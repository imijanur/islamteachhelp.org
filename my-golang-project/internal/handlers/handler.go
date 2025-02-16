package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/imijanur/islamteachhelps.com/internal/models"
)

var footer = models.Footer{
	EshtablishedYear:     1972,
	EshtablishedYearUrdu: 1392,
	CopyRightStartYear:   2023,
	Mobile:               "+91 93305 61238",
	PrimaryEmail:         "bmkuloom@gmail.com",
	SecondaryEmail:       "contact@islamteachhelp.org",
	Address:              "Vill: Baladevpur, P.O: Ramakantanagar, P.S: Mandirbazar, Dist: South 24 Parganas, West Bengal, India, Pin: 743395",
	AddressUrdu:          "گاؤں: بلادیوپور، پو: راماکنتنگر، پی ایس: مندربازار، ضلع: جنوبی 24 پرگناز، مغربی بنگال، بھارت، پن: 743395",
	About:                "The Madrasah was established in 1972. The Madrasah is situated in the village of Baladevpur, P.O: Ramakantanagar, P.S: Mandirbazar, Dist: South 24 Parganas, West Bengal, India, Pin: 743395. The Madrasah is a registered institution under the West Bengal Society Registration Act, 1961. The Madrasah is also registered under the Foreign Contribution Regulation Act, 2010 (FCRA). The Madrasah is also registered under the Income Tax Act, 1961 under section 12A and 80G. The Madrasah is also registered under the National Institute for Empowerment of Persons with Multiple Disabilities (NIEPMD) under the Department of Empowerment of Persons with Disabilities (Divyangjan), Ministry of Social Justice and Empowerment, Government of India. The Madrasah is also registered under the National Trust for the Welfare of Persons with Autism, Cerebral Palsy, Mental Retardation and Multiple Disabilities, Ministry of Social Justice and Empowerment, Government of India.",
	AboutUrdu:            "مدرسہ 1972 میں قائم کیا گیا تھا۔ مدرسہ بلادیوپور گاؤں میں واقع ہے، پو: راماکنتنگر، پی ایس: مندربازار، ضلع: جنوبی 24 پرگناز، مغربی بنگال، بھارت، پن: 743395۔ مدرسہ مغربی بنگال سوسائٹی رجسٹریشن ایکٹ، 1961 کے تحت رجسٹر شدہ ادارہ ہے۔ مدرسہ فارن کنٹری بحران ایکٹ، 2010 (ایف سی آر اے) کے تحت بھی رجسٹر شدہ ہے۔ مدرسہ انکم ٹیکس ایکٹ، 1961 کے تحت سیکشن 12A اور 80G کے تحت بھی رجسٹر شدہ ہے۔ مدرسہ انکم ٹیکس ایکٹ، 1961 کے تحت سیکشن 12A اور 80G کے تحت بھی رجسٹر شدہ ہے۔ مدرسہ انکم ٹیکس ایکٹ، 1961 کے تحت سیکشن 12A اور 80G کے تحت بھی رجسٹر شدہ ہے۔ مدرسہ انکم ٹیکس ا",
}

func HomeHandler(c fiber.Ctx) error {
	return renderTemplate(c, "index", models.Page{
		Title:  "Home Page",
		Footer: footer,
	})
}

func TestHandler(c fiber.Ctx) error {
	return renderTemplate(c, "test", models.Page{
		Title:  "Test Page",
		Footer: footer,
	})
}
func AboutHandler(c fiber.Ctx) error {
	return renderTemplate(c, "about", models.Page{
		Title:  "ABOUT &#8211; Donation",
		Footer: footer,
	})
}
func ActivitiesHandler(c fiber.Ctx) error {
	data := models.ActivityPageData{
		Activities: []models.Activity{
			{
				ID:              "1",
				Description:     "Giving sacrificial animals to the villages of the poor at the time of Qurbani",
				DescriptionUrdu: "قربانی کے وقت غریبوں کے دیہاتوں میں قربانی کے جانور دینا",
				Images:          []string{"/static/img/activities/Kurbani-1.jpg"},
			},
			{
				ID:              "2",
				Description:     "To help those who cannot marry their daughters",
				DescriptionUrdu: "ان لوگوں کی مدد کرنا جو اپنی بیٹیوں کی شادی نہیں کر سکتے",
				Images:          []string{"/static/img/activities/rice-distribution-1.jpg"},
			},
			{
				ID:              "3",
				Description:     "Issuance of qarje hasana for business",
				DescriptionUrdu: "کاروبار کے لیے قرض حسنہ کا اجرا",
				Images:          []string{"/static/img/activities/rice-distribution-2.jpg"},
			},
			{
				ID:              "4",
				Description:     "Distributing Iftar among the poor during Fasting",
				DescriptionUrdu: "روزے کے دوران غریبوں میں افطار تقسیم کرنا",
				Images:          []string{"/static/img/activities/iftar-1.jpg", "/static/img/activities/iftar-2.jpg"},
			},
			{
				ID:              "5",
				Description:     "Widows should not beg, so we buy them sewing machines and goats",
				DescriptionUrdu: "بیواؤں کو بھیک نہیں مانگنی چاہیے، اس لیے ہم انہیں سلائی مشینیں اور بکریاں خرید کر دیتے ہیں",
				Images:          []string{"/static/img/activities/rice-distribution-3.jpg"},
			},
			{
				ID:              "6",
				Description:     "Buying education materials for poor children",
				DescriptionUrdu: "غریب بچوں کے لیے تعلیمی مواد خریدنا",
				Images:          []string{"/static/img/activities/treatment-1.jpg"},
			},
			{
				ID:              "7",
				Description:     "Buying vans for poor families",
				DescriptionUrdu: "غریب خاندانوں کے لیے وین خریدنا",
				Images:          []string{"/static/img/activities/treatment-2.jpg"},
			},
			{
				ID:              "8",
				Description:     "Together the marriages of poor people's daughters are arranged and then their expenses are borne",
				DescriptionUrdu: "غریب لوگوں کی بیٹیوں کی شادیاں مل کر کروائی جاتی ہیں اور پھر ان کے اخراجات برداشت کیے جاتے ہیں",
				Images:          []string{"/static/img/activities/Tube-well-1.jpg", "/static/img/activities/tube-well-2.jpg", "/static/img/activities/tube-well-3.jpg", "/static/img/activities/tube-well-4.jpg", "/static/img/activities/tube-well-5.jpg", "/static/img/activities/tube-well-6.jpg", "/static/img/activities/tube-well-7.jpg", "/static/img/activities/tube-well-with-motor.jpg", "/static/img/activities/water-tank-1.jpg"},
			},
		},
	}

	return renderTemplate(c, "activities", models.Page{
		Title:  "ACTIVITIES &#8211; Donation",
		Footer: footer,
		Data:   data,
	})
}
func AuthorHandler(c fiber.Ctx) error {
	return renderTemplate(c, "author", models.Page{
		Title:  "Author Page",
		Footer: footer,
	})
}
func CategoryHandler(c fiber.Ctx) error {
	return renderTemplate(c, "category", models.Page{
		Title:  "Uncategorized &#8211; Donation",
		Footer: footer,
	})
}
func CommentsHandler(c fiber.Ctx) error {
	return renderTemplate(c, "test", models.Page{
		Title:  "Test Page",
		Footer: footer,
	})
}
func ContactHandler(c fiber.Ctx) error {
	contactDetails := models.ContactDetails{
		Title: map[string]string{
			"en": "Contact Us",
			"ur": "ہم سے رابطہ کریں",
		},
		Address: map[string]string{
			"en": footer.Address,
			"ur": footer.AddressUrdu,
		},
		PrimaryEmail:   footer.PrimaryEmail,
		SecondaryEmail: footer.SecondaryEmail,
		Phone:          footer.Mobile,
		FacebookURL:    "https://facebook.com",
		TwitterURL:     "https://twitter.com",
		YouTubeURL:     "https://youtube.com",
	}

	return renderTemplate(c, "contact", models.Page{
		Title:  "CONTACT &#8211; Donation",
		Footer: footer,
		Data:   contactDetails,
	})
}
func DonationHandler(c fiber.Ctx) error {
	return renderTemplate(c, "donation", models.Page{
		Title:  "Donation &#8211; Donation",
		Footer: footer,
	})
}
func FeedHandler(c fiber.Ctx) error {
	return renderTemplate(c, "feed", models.Page{
		Title:  "Feed &#8211; Donation",
		Footer: footer,
	})
}
func GalleryHandler(c fiber.Ctx) error {
	return renderTemplate(c, "gallery", models.Page{
		Title:  "GALLERY &#8211; Donation",
		Footer: footer,
	})
}
func MadrashaHandler(c fiber.Ctx) error {
	return renderTemplate(c, "madrasha", models.Page{
		Title:  "Madrasha &#8211; Donation",
		Footer: footer,
	})
}
func BalanceSheetHandler(c fiber.Ctx) error {
	return renderTemplate(c, "balance-sheet", models.Page{
		Title:  "Balance Sheet",
		Footer: footer,
	})
}
func SamplePageHandler(c fiber.Ctx) error {
	return renderTemplate(c, "sample-page", models.Page{
		Title:  "Sample Page &#8211; Donation",
		Footer: footer,
	})
}
func South24ParganasDistrictHandler(c fiber.Ctx) error {
	return renderTemplate(c, "south-24-parganas-district", models.Page{
		Title:  "South 24 Parganas District &#8211; Donation",
		Footer: footer,
	})
}

func NotFoundHandler(c fiber.Ctx) error {
	return renderTemplate(c, "notfound", models.Page{
		Title:  "Not Found",
		Footer: footer,
	})
}
