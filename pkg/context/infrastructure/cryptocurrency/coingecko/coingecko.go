package coingecko

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/bastean/tgo/pkg/context/domain/errors"
)

type CoinGecko struct {
	key    string
	url    *url.URL
	client *http.Client
}

func (api *CoinGecko) value(values ...string) []string {
	return []string{strings.Join(values, ",")}
}

func (api *CoinGecko) request(endpoint string, values *url.Values) ([]byte, error) {
	url := api.url.JoinPath(endpoint)

	if values != nil {
		url.RawQuery = values.Encode()
	}

	request, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "request",
			What:  "Failure to create a new request",
			Why: errors.Meta{
				"URL": url.String(),
			},
			Who: err,
		})
	}

	request.Header.Add("accept", "application/json")
	request.Header.Add("x-cg-demo-api-key", api.key)

	response, err := api.client.Do(request)

	switch {
	case err != nil:
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "request",
			What:  "Failure to send a request",
			Why: errors.Meta{
				"URL": url.String(),
			},
			Who: err,
		})
	case response.StatusCode != http.StatusOK:
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "request",
			What:  "Failure request",
			Why: errors.Meta{
				"URL":    url.String(),
				"Status": response.Status,
			},
		})
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "request",
			What:  "Failure to read the response body",
			Why: errors.Meta{
				"URL": url.String(),
			},
			Who: err,
		})
	}

	return body, nil
}

func (api *CoinGecko) decode(data []byte, target any) error {
	err := json.Unmarshal(data, &target)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Decode",
			What:  "Failure to JSON decoding",
			Who:   err,
		})
	}

	return nil
}

func (api *CoinGecko) Status() error {
	_, err := api.request("ping", nil)

	if err != nil {
		return errors.BubbleUp(err, "Status")
	}

	return nil
}

func (api *CoinGecko) IsCurrencySupported(currency string) error {
	currency = strings.ToLower(currency)

	response, err := api.request("simple/supported_vs_currencies", nil)

	if err != nil {
		return errors.BubbleUp(err, "Status")
	}

	var currencies []string

	err = api.decode(response, &currencies)

	if err != nil {
		return errors.BubbleUp(err, "Status")
	}

	if !slices.Contains(currencies, currency) {
		return errors.NewFailure(&errors.Bubble{
			Where: "IsCurrencySupported",
			What:  "Unsupported currency",
			Why: errors.Meta{
				"Currency": currency,
			},
		})
	}

	return nil
}

func (api *CoinGecko) CoinPrices(currency string, coins []string) (map[string]float64, error) {
	currency = strings.ToLower(currency)

	response, err := api.request("simple/price", &url.Values{
		"ids":           api.value(coins...),
		"vs_currencies": api.value(currency),
		"precision":     api.value("8"),
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "CoinPrices")
	}

	var data map[string]map[string]float64

	err = api.decode(response, &data)

	if err != nil {
		return nil, errors.BubbleUp(err, "CoinPrices")
	}

	prices := map[string]float64{}

	var price float64

	for _, coin := range coins {
		price = data[coin][currency]

		if price >= 1 {
			price = math.Floor(price*100) / 100
		}

		prices[coin] = price
	}

	return prices, nil
}

func New(key, rawURL string) (*CoinGecko, error) {
	coingecko := &CoinGecko{
		key: key,
		client: &http.Client{
			Timeout: time.Minute,
		},
	}

	url, err := url.Parse(rawURL)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "New",
			What:  "Failure to parse URL",
			Why: errors.Meta{
				"URL": rawURL,
			},
			Who: err,
		})
	}

	coingecko.url = url

	err = coingecko.Status()

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	return coingecko, nil
}
