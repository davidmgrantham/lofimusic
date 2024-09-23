package main

import (
	"path"
	"sort"
	"strings"
	"time"
)

type liveRadio struct {
	Slug    string
	Name    string
	Owner   string
	URL     string
	Cards   []string
	Links   []socialLink
	AddedAt time.Time
}

func (r liveRadio) youtubeID() string {
	return path.Base(r.URL)
}

func getLiveRadios() []liveRadio {
	radios := []liveRadio{
		{
			Slug:  "chillhop",
			Name:  "Chillhop Raccoon",
			Owner: "Chillhop Music",
			URL:   "https://youtu.be/5yx6BWlEVcY",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/5yx6BWlEVcY",
				},
			},
		},
		{
			Slug:  "chillhop-relax",
			Name:  "Chillhop Relaxing Raccoon",
			Owner: "Chillhop Music",
			URL:   "https://youtu.be/7NOSDKb0HlU",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/7NOSDKb0HlU",
				},
			},
		},
		{
			Slug:  "taiki",
			Name:  "Taiki",
			Owner: "Chill with Taiki",
			URL:   "https://youtu.be/qH3fETPsqXU",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/qH3fETPsqXU",
				},
			},
		},
		{
			Slug:  "fantasy-lofi",
			Name:  "Fantasy Lofi",
			Owner: "FantasyLofi",
			URL:   "https://youtu.be/WtEJFKg-LtE",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/WtEJFKg-LtE",
				},
			},
		},
		{
			Slug:  "tavern-music",
			Name:  "Tavern Music",
			Owner: "Filip Lackovic",
			URL:   "https://youtu.be/vK5VwVyxkbI",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/vK5VwVyxkbI",
				},
			},
		},
		{
			Slug:  "video-game-lofi",
			Name:  "Video Game Lofi",
			Owner: "GlitchxCity",
			URL:   "https://youtu.be/akW7qWS_p-g",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/akW7qWS_p-g",
				},
			},
		},
		{
			Slug:  "lofigirl",
			Name:  "Lofi Girl",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/jfKfPfyJRdk",
			Cards: []string{
				"Lofi girl is a radio that broadcasts lo-fi hip hop songs created by a French fellow named Dimitri in 2017.",
				`The animation, made by Juan Pablo Machado, is modeled after Shizuku Tsukishima, a girl character from the Studio Ghibli film "Whisper of the Heart".`,
				"Named Jade, the Lofi girl is shown studying in Lyon, a city from France where her designer Juan Pablo used to live.",
				"The view through the window depicts the buildings on the slopes of Croix-Rousse, where the bell tower of the Bon-Pasteur church can be spotted.",
			},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/jfKfPfyJRdk",
				},
			},
		},
		{
			Slug:  "ultima-online-lofi",
			Name:  "Ultima Online Lofi",
			Owner: "LootGoblinLofi",
			URL:   "https://youtu.be/c4Tb4RlcUtE",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/c4Tb4RlcUtE",
				},
			},
		},
		{
			Slug:  "everything-fades-to-blue",
			Name:  "Everything Fades To Blue",
			Owner: "Sleepy Fish",
			URL:   "https://youtu.be/PfgS405CdXk",
			Cards: []string{
				"Everything Fades To Blue is a mix of indie/emo and ambient music produced by Sleep Fish, a Pennsylvania-based producer that is also a student in statistics, data science, and machine learning.",
				"Sleepy Fish is one of the few Lo-fi artists who actually sing in its creations.",
				"Everything Fades To Blue is the 3rd episode of a story where a tidal wave destroys an island along with the home where Sleepy Fish used to live.",
				"Toppled into the sea, on its own for the first time, Sleepy Fish uses its glow to search for family, to guide others, and to find its way.",
				`The episode comes after "My Room Becomes the Sea" and "Beneath Your Waves".`,
				"The undersea-themed animation has been made in collaboration with Tristan Gion and Bien à Vous Studio, a French studio based in Nantes.",
			},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://youtu.be/PfgS405CdXk",
				},
			},
		},
		{
			Slug:  "medievallofi",
			Name:  "Medieval Lofi",
			Owner: "Medieval Lofi",
			URL:   "https://www.youtu.be/4vmrqCrImSE",
			Cards: []string{},
			Links: []socialLink{
				{
					Slug: "youtube",
					URL:  "https://www.youtu.be/4vmrqCrImSE",
				},
			},
		},
	}

	sort.Slice(radios, func(a, b int) bool {
		return strings.Compare(radios[a].Name, radios[b].Name) < 0
	})

	for _, r := range radios {
		sort.Slice(r.Links, func(a, b int) bool {
			return strings.Compare(r.Links[a].Slug, r.Links[b].Slug) < 0
		})
	}

	return radios
}

type socialLink struct {
	Slug string
	URL  string
}

func socialIcon(slug string) string {
	switch slug {
	case "youtube":
		return youtubeSVG

	case "reddit":
		return redditSVG

	case "facebook":
		return facebookSVG

	case "instagram":
		return instagramSVG

	case "twitter":
		return twitterSVG

	case "spotify":
		return spotifySVG

	case "discord":
		return discordSVG

	case "website":
		return websiteSVG

	default:
		return linkSVG
	}
}
