/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package data

import "time"

// CachedSubscription is the entity that we use to store the cached
// Stripe subscription API call results. We make a request to Stripe
// API and cache the response IF there is a match. If there is no
// matching subscription, we don’t cache the response. The TTL for the
// cache is 24 hours.
type CachedSubscription struct {
	Email               string
	StripeProductId     string
	StripePriceId       string
	StripeCurrency      string
	StripeUnitAmount    int
	StripeInterval      string // "month"
	StriveIntervalCount int    // 1: monthly, 12: yearly
	RecordCreated       time.Time
}
