/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package data

type Subscription struct {
	StripeProductId     string
	StripePriceId       string
	StripeCurrency      string
	StripeUnitAmount    int
	StripeInterval      string // "month"
	StriveIntervalCount int    // 1: monthly, 12: yearly
}
