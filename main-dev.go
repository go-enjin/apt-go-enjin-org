//go:build dev || !prd

// Copyright (c) 2023  The Go-Enjin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	apt "github.com/go-enjin/apt-enjin-theme"
	semantic "github.com/go-enjin/semantic-enjin-theme"

	"github.com/go-enjin/be/features/fs/content"
	"github.com/go-enjin/be/features/fs/menu"
	"github.com/go-enjin/be/features/fs/public"
	"github.com/go-enjin/be/features/fs/themes"
)

func init() {
	fThemes = themes.New().
		Include(semantic.Theme()).
		Include(apt.Theme()).
		SetTheme(apt.Name).
		Make()

	fMenus = menu.New().
		MountLocalPath("menus", "menus").
		Make()

	fPublic = public.New().
		MountLocalPath("/", "public").
		Make()

	fAptRepo = public.NewTagged("fs-public-apt-repo").
		MountLocalPath("/"+UseAptFlavour, UseBasePath+"/"+UseAptFlavour).
		SetRegexCacheControl("/dists/", "no-store").
		Make()

	fContent = content.New().
		MountLocalPath("/", "content").
		AddToIndexProviders("pages-pql").
		AddToSearchProviders("bleve-fts").
		Make()
}