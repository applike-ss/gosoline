package adwords

import "time"

type AdDimensions struct {
	Width  uint
	Height uint
}
type MediaSizeDimensionsMapEntry struct {
	Key   string
	Value AdDimensions
}
type MediaSizeStringMapEntry struct {
	Key   string
	Value string
}

type AdMedia struct {
	CreationTime *time.Time
	Dimensions   []*MediaSizeDimensionsMapEntry
	FileSize     *uint
	MediaId      *uint
	MediaType    *string `xml:"Media.Type"`
	MimeType     *string
	Name         *string
	ReferenceId  *uint
	SourceUrl    *string
	Type         *string `xml:"type"`
	Urls         []*MediaSizeStringMapEntry
}

type Media struct {
	AdMedia
	Data string `xml:"data"`
}

// OS_TYPE_IOS/OS_TYPE_ANDROID/UNKNOWN
type OsType string

type AppUrl struct {
	Url    string `xml:"url"`
	OsType OsType `xml:"osType"`
}

type CustomParameter struct {
	Key      string
	Value    string
	IsRemove bool
}

type CustomParameters struct {
	Parameters []CustomParameter
	DoReplace  bool
}

type UrlList struct {
	Urls []string
}

type UrlData struct {
	UrlId               string
	FinalUrls           UrlList
	FinalMobileUrls     UrlList
	TrackingUrlTemplate string
}

// DEPRECATED_AD/IMAGE_AD/PRODUCT_AD/TEMPLATE_AD/TEXT_AD/THIRD_PARTY_REDIRECT_AD/DYNAMIC_SEARCH_AD/CALL_ONLY_AD/EXPANDED_TEXT_AD/RESPONSIVE_DISPLAY_AD/SHOWCASE_AD/GOAL_OPTIMIZED_SHOPPING_AD/EXPANDED_DYNAMIC_SEARCH_AD/GMAIL_AD/RESPONSIVE_SEARCH_AD/MULTI_ASSET_RESPONSIVE_DISPLAY_AD/UNIVERSAL_APP_AD/UNKNOWN
type AdType string

// UNKNOWN/AD_VARIATIONS
type SystemManagedEntitySource string

type AdImageReponse struct {
	Id                        uint                      `xml:"id"` // This should be long, but golang doesn't know long :-/
	URL                       string                    `xml:"url"`
	DisplayUrl                string                    `xml:"displayUrl"`
	FinalUrls                 []string                  `xml:"finalUrls"`
	FinalMobileUrls           []string                  `xml:"finalMobileUrls"`
	FinalAppUrls              []AppUrl                  `xml:"finalAppUrls"`
	TrackingUrlTemplate       string                    `xml:"trackingUrlTemplate"`
	FinalUrlSuffix            string                    `xml:"finalUrlSuffix"`
	UrlCustomParameters       CustomParameters          `xml:"urlCustomParameters"`
	UrlData                   UrlData                   `xml:"urlData"`
	Automated                 bool                      `xml:"automated"`
	Type                      AdType                    `xml:"type"`
	DevicePreference          uint                      `xml:"devicePreference"` // This should be long, but golang doesn't know long :-/
	SystemManagedEntitySource SystemManagedEntitySource `xml:"systemManagedEntitySource"`
	AdType                    string                    `xml:"Ad.Type"`
	Image                     string                    `xml:"image"`
	Name                      string                    `xml:"name"`
	AdToCopyImageFrom         uint                      `xml:"adToCopyImageFrom"` // This should be long, but golang doesn't know long :-/
}
