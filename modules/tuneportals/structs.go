package tuneportals

import (
	"tunebot/util/proxies"

	"net/http"
)

type Task struct {
	Site           string
	UPC            string
	Quantity       string
	ShippingOption string
	PriceLimit     float64

	Stage    string
	Callback string

	Count    int
	UseProxy bool
	Proxies  *proxies.ProxyList
	Time     int64
	Webhook  string

	FirstName    string
	LastName     string
	Email        string
	Phone        string
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Zip          string
	Country      string
	CardNumber   string
	CardMonth    string
	CardYear     string
	CVC          string

	image     string
	name      string
	price     string
	mediaType string

	stripeKey        string
	itemId           int64
	itemType         string
	total            string
	paymentToken     string
	sessionTimestamp int64

	client *http.Client
}

const (
	INIT             = "init"
	PRELOAD_CHECKOUT = "preload_checkout"
	CREATE_SESSION   = "create_session"
	LOAD_PRODUCT     = "load_product"
	ADD_TO_CART      = "add_to_cart"
	LOAD_CART        = "load_cart"
	LOAD_SHIPPING    = "load_shipping"
	TOKENIZE_PAYMENT = "tokenize_payment"
	SUBMIT_ORDER     = "submit_order"
	KILL             = "kill"
)

type productResponse struct {
	Success int `json:"success"`
	Item    struct {
		SimilarItemLoop []struct {
			PreorderFlag               int         `json:"preorder_flag"`
			NewInventoryAvailable      interface{} `json:"new_inventory_available"`
			ItemID                     int64       `json:"item_id"`
			NewLocalInventoryAvailable interface{} `json:"new_local_inventory_available"`
			UsedPrice                  interface{} `json:"used_price"`
			HasStock                   interface{} `json:"has_stock"`
			Owner                      string      `json:"owner"`
			ItemURL                    string      `json:"item_url"`
			UsedCount                  interface{} `json:"used_count"`
			Digital                    string      `json:"digital"`
			Format                     string      `json:"format"`
			Price                      interface{} `json:"price"`
			Title                      string      `json:"title"`
			UsedOnly                   interface{} `json:"used_only"`
			MnetID                     string      `json:"mnet_id"`
			LowestUsedPrice            interface{} `json:"lowest_used_price"`
		} `json:"similar_item_loop"`
		RegularPrice      string        `json:"regular_price"`
		HasTracklist      int           `json:"has_tracklist"`
		BroadtimeEncoded  int           `json:"broadtime_encoded"`
		ItemOwner         string        `json:"item_owner"`
		BlurbText         string        `json:"blurb_text"`
		ReviewAuthor      string        `json:"review_author"`
		TypeID            int           `json:"type_id"`
		ExemptTax         string        `json:"exempt_tax"`
		ArtistID          int           `json:"artist_id"`
		GenreLoop         interface{}   `json:"genre_loop"`
		StrippedBlurbText string        `json:"stripped_blurb_text"`
		Suppliers         []interface{} `json:"suppliers"`
		RelatedItemIds    []int64       `json:"related_item_ids"`
		Details           string        `json:"details"`
		OwnerTypeID       int           `json:"owner_type_id"`
		HTMLTitle         string        `json:"html_title"`
		GenreID           interface{}   `json:"genre_id"`
		ProductSku        interface{}   `json:"product_sku"`
		ItemDetails       string        `json:"item_details"`
		FeatureLoop       []struct {
			FeatureID   int         `json:"feature_id"`
			SiteID      interface{} `json:"site_id"`
			LockedFlag  interface{} `json:"locked_flag"`
			GenreID     int         `json:"genre_id"`
			OwnerID     int         `json:"owner_id"`
			SiteName    interface{} `json:"site_name"`
			SpecialFlag interface{} `json:"special_flag"`
			GenreName   string      `json:"genre_name"`
		} `json:"feature_loop"`
		ExemptShipping string      `json:"exempt_shipping"`
		MissingGraphic interface{} `json:"missing_graphic"`
		Title          string      `json:"title"`
		KioskItemURL   string      `json:"kiosk_item_url"`
		InventoryList  []struct {
			ConditionID             int         `json:"condition_id"`
			MaxPerCart              interface{} `json:"max_per_cart"`
			ExemptTax               string      `json:"exempt_tax"`
			SaleExpireDays          int         `json:"sale_expire_days"`
			InventoryID             int         `json:"inventory_id"`
			ExemptPriceRule         string      `json:"exempt_price_rule"`
			ID                      int         `json:"id"`
			Updated                 string      `json:"updated"`
			PreorderLimitByQuantity string      `json:"preorder_limit_by_quantity"`
			SaleExpireDate          interface{} `json:"sale_expire_date"`
			FixedPrice              string      `json:"fixed_price"`
			Owner                   string      `json:"owner"`
			PreorderBy              string      `json:"preorder_by"`
			Cost                    string      `json:"cost"`
			StatusCode              string      `json:"status_code"`
			Protected               string      `json:"protected"`
			UpdateDtime             interface{} `json:"update_dtime"`
			Quantity                int         `json:"quantity"`
			StoreNote               interface{} `json:"store_note"`
			SaleUntilDays           int         `json:"sale_until_days"`
			SalePrice               interface{} `json:"sale_price"`
			InsertDtime             interface{} `json:"insert_dtime"`
			OwnerTypeID             int         `json:"owner_type_id"`
			SaleStartDate           interface{} `json:"sale_start_date"`
			PreorderBlock           string      `json:"preorder_block"`
			ProductSku              string      `json:"product_sku"`
			Notes                   string      `json:"notes"`
			ExemptShipping          string      `json:"exempt_shipping"`
			OwnerID                 int         `json:"owner_id"`
			Condition               string      `json:"condition"`
		} `json:"inventory_list"`
		MnetTpid              int         `json:"mnet_tpid"`
		PrettyDate            string      `json:"pretty_date"`
		FutureRelease         int         `json:"future_release"`
		VinylSurcharge        int         `json:"vinyl_surcharge"`
		HasSamples            int         `json:"has_samples"`
		SamplexmlURL          string      `json:"samplexml_url"`
		NewInventoryCount     int         `json:"new_inventory_count"`
		MediaType             string      `json:"media_type"`
		ItemTitle             string      `json:"item_title"`
		DigitalOnly           interface{} `json:"digital_only"`
		NewPrice              string      `json:"new_price"`
		ItemURL               string      `json:"item_url"`
		StatusCode            string      `json:"status_code"`
		HasLivePlayart        string      `json:"has_live_playart"`
		UpdateDtime           int         `json:"update_dtime"`
		MnetExactMatch        int         `json:"mnet_exact_match"`
		WeOwnThis             int         `json:"we_own_this"`
		CoverArtFlag          int         `json:"cover_art_flag"`
		VendorSku             interface{} `json:"vendor_sku"`
		GoLiveDatetime        string      `json:"go_live_datetime"`
		MnetID                string      `json:"mnet_id"`
		ImageServerFlag       int         `json:"image_server_flag"`
		PreorderFlag          int         `json:"preorder_flag"`
		Review                string      `json:"review"`
		RelatedItemsBlob      string      `json:"related_items_blob"`
		CatalogNumber         string      `json:"catalog_number"`
		LockReleaseDate       string      `json:"lock_release_date"`
		NewInventoryAvailable int         `json:"new_inventory_available"`
		MaxPerCart            interface{} `json:"max_per_cart"`
		ItemID                int64       `json:"item_id"`
		PhotoLoop             []struct {
			Source      string `json:"source"`
			FileSize    int    `json:"file_size"`
			PhotoID     int64  `json:"photo_id"`
			UpdateDtime int    `json:"update_dtime"`
			InsertDtime int    `json:"insert_dtime"`
			OwnerTypeID int    `json:"owner_type_id"`
			FileType    string `json:"file_type"`
			Caption     string `json:"caption"`
			Title       string `json:"title"`
			OwnerID     int    `json:"owner_id"`
		} `json:"photo_loop"`
		ArtistLoop []struct {
			LockNameFlag string `json:"lock_name_flag"`
			Primary      int    `json:"primary"`
			Name         string `json:"name"`
			OriginalName string `json:"original_name"`
			UpdateDtime  int    `json:"update_dtime"`
			InsertDtime  int    `json:"insert_dtime"`
			OwnerTypeID  int    `json:"owner_type_id"`
			ArtistID     int    `json:"artist_id"`
			OwnerID      int    `json:"owner_id"`
			MnetID       int    `json:"mnet_id"`
			WebsiteURL   string `json:"website_url"`
		} `json:"artist_loop"`
		ItemArtist             string      `json:"item_artist"`
		Cost                   string      `json:"cost"`
		PreorderInventoryCount interface{} `json:"preorder_inventory_count"`
		PhotoID                int64       `json:"photo_id"`
		MetaDescription        string      `json:"meta_description"`
		ReleaseDate            string      `json:"release_date"`
		ArtistName             string      `json:"artist_name"`
		OrderingBlocked        string      `json:"ordering_blocked"`
		ThumbnailURL           string      `json:"thumbnail_url"`
		Msrp                   string      `json:"msrp"`
		Genre                  interface{} `json:"genre"`
		MnetPrice              string      `json:"mnet_price"`
		EncodedItemID          int         `json:"encoded_item_id"`
		ArtistURL              string      `json:"artist_url"`
		Label                  string      `json:"label"`
		Price                  string      `json:"price"`
		GroupMe                string      `json:"group_me"`
		Upc                    string      `json:"upc"`
		Status                 interface{} `json:"status"`
		PhotoOwnerID           int         `json:"photo_owner_id"`
		ExemptPriceRule        string      `json:"exempt_price_rule"`
		ItemOwnerID            int         `json:"item_owner_id"`
		InPrint                int         `json:"in_print"`
		UsedInventoryCount     int         `json:"used_inventory_count"`
		DefaultWeight          string      `json:"default_weight"`
		InsertDtime            int         `json:"insert_dtime"`
		UpdateSolr             string      `json:"update_solr"`
		CanonicalURL           string      `json:"canonical_url"`
		Artist                 string      `json:"artist"`
		BlurbID                int         `json:"blurb_id"`
		FormatIcon             interface{} `json:"format_icon"`
		CoverURL               string      `json:"cover_url"`
		Format                 string      `json:"format"`
		ArtistPhoto            string      `json:"artist_photo"`
		DiscLoop               []struct {
			TrackLoop []struct {
				TrackName      string `json:"track_name"`
				TrackNumber    string `json:"track_number"`
				HasSample      int    `json:"has_sample"`
				MnetTpid       int    `json:"mnet_tpid"`
				TrackID        int    `json:"track_id"`
				MnetTrackID    string `json:"mnet_track_id"`
				MnetPrice      string `json:"mnet_price"`
				Index          int    `json:"index"`
				MnetTrackPrice string `json:"mnet_track_price"`
				MnetTrackTpid  int    `json:"mnet_track_tpid"`
				DiscNumber     string `json:"disc_number"`
				MnetID         string `json:"mnet_id"`
			} `json:"track_loop"`
			TrackCount int    `json:"track_count"`
			DiscID     int    `json:"disc_id"`
			DiscNumber string `json:"disc_number"`
		} `json:"disc_loop"`
		OwnerID          int    `json:"owner_id"`
		DisplayPrice     string `json:"display_price"`
		KioskFullPreview string `json:"kiosk_full_preview"`
		Vendor           string `json:"vendor"`
	} `json:"item"`
}

type atcResponse struct {
	AnonCart []struct {
		CartItemQuantity string `json:"cart_item_quantity"`
		CartItemID       string `json:"cart_item_id"`
		CartItemType     string `json:"cart_item_type"`
	} `json:"anon_cart"`
	Session string `json:"session"`
	Message string `json:"message"`
}

type cartResponse struct {
	Taxes                string      `json:"taxes"`
	AdjustedSubtotal     interface{} `json:"adjusted_subtotal"`
	Subtotal             string      `json:"subtotal"`
	DiscountableSubtotal string      `json:"discountable_subtotal"`
	Total                string      `json:"total"`
	Digital              interface{} `json:"digital"`
	Fees                 interface{} `json:"fees"`
	Shipping             interface{} `json:"shipping"`
	Physical             string      `json:"physical"`
	Prices               struct {
		Taxes                string      `json:"taxes"`
		AdjustedSubtotal     interface{} `json:"adjusted_subtotal"`
		Subtotal             string      `json:"subtotal"`
		DiscountableSubtotal string      `json:"discountable_subtotal"`
		Total                string      `json:"total"`
		Digital              interface{} `json:"digital"`
		Fees                 interface{} `json:"fees"`
		Physical             string      `json:"physical"`
		Shipping             interface{} `json:"shipping"`
		CouponMessage        interface{} `json:"coupon_message"`
	} `json:"prices"`
	CouponMessage interface{} `json:"coupon_message"`
}

type shippingResponse struct {
	Errors          []interface{} `json:"errors"`
	ShippingOptions []struct {
		Provider        string `json:"provider"`
		AecCode         string `json:"aec_code"`
		AdditionalVinyl string `json:"additional_vinyl"`
		AdditionalItem  string `json:"additional_item"`
		FirstItem       string `json:"first_item"`
		CartMinimum     string `json:"cart_minimum"`
		ID              int    `json:"id"`
		International   string `json:"international"`
		DeliveryType    string `json:"delivery_type"`
	} `json:"shipping_options"`
}

type stripeResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Card   struct {
		ID                 string      `json:"id"`
		Object             string      `json:"object"`
		AddressCity        interface{} `json:"address_city"`
		AddressCountry     interface{} `json:"address_country"`
		AddressLine1       interface{} `json:"address_line1"`
		AddressLine1Check  interface{} `json:"address_line1_check"`
		AddressLine2       interface{} `json:"address_line2"`
		AddressState       interface{} `json:"address_state"`
		AddressZip         string      `json:"address_zip"`
		AddressZipCheck    string      `json:"address_zip_check"`
		Brand              string      `json:"brand"`
		Country            string      `json:"country"`
		CvcCheck           string      `json:"cvc_check"`
		DynamicLast4       interface{} `json:"dynamic_last4"`
		ExpMonth           int         `json:"exp_month"`
		ExpYear            int         `json:"exp_year"`
		Funding            string      `json:"funding"`
		Last4              string      `json:"last4"`
		Name               interface{} `json:"name"`
		TokenizationMethod interface{} `json:"tokenization_method"`
		Wallet             interface{} `json:"wallet"`
	} `json:"card"`
	ClientIP string `json:"client_ip"`
	Created  int    `json:"created"`
	Livemode bool   `json:"livemode"`
	Type     string `json:"type"`
	Used     bool   `json:"used"`
}

type orderResponse struct {
	ForceCartRefresh int         `json:"force_cart_refresh"`
	Errors           []string    `json:"errors"`
	PhysicalComplete int         `json:"physical_complete"`
	DigitalComplete  int         `json:"digital_complete"`
	RedirectURL      string      `json:"redirect_url"`
	ForceLogin       int         `json:"force_login"`
	MobileUpdates    interface{} `json:"mobile_updates"`
	IncompleteUser   int         `json:"incomplete_user"`
}
