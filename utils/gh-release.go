/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package utils

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	sveltinerr "github.com/sveltinio/sveltin/internal/errors"
	"github.com/tidwall/gjson"
)

// RetrieveLatestReleaseVersion returns the latest available sveltin release version as string.
func RetrieveLatestReleaseVersion() (string, error) {
	var version string
	url := "https://api.github.com/repos/sveltinio/sveltin/releases/latest"

	httpClient := http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", sveltinerr.NewCheckReleaseError(url)
	}

	response, err := httpClient.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return "", sveltinerr.NewDefaultError(err)
		}
		return "", sveltinerr.NewCheckReleaseError(url)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", sveltinerr.NewParseResponseError(err)
	}

	value := gjson.Get(string(data), "tag_name")
	if value.Exists() {
		version = strings.Replace(value.String(), "v", "", 1)
	}

	return version, nil
}
