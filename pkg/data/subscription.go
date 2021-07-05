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
	Email               string    `json:"email" bson:"email"`
	StripeProductId     string    `json:"stripeProductId" bson:"stripeProductId"`
	StripePriceId       string    `json:"stripePriceId" bson:"stripePriceId"`
	StripeCurrency      string    `json:"stripeCurrency" bson:"stripeCurrency"`
	StripeUnitAmount    int       `json:"stripeUnitAmount" bson:"stripeUnitAmount"`
	StripeInterval      string    `json:"stripeInterval" bson:"stripeInterval"`           // "month"
	StripeIntervalCount int       `json:"stripeIntervalCount" bson:"stripeIntervalCount"` // 1: monthly, 12: yearly
	RecordCreated       time.Time `json:"recordCreated" bson:"recordCreated"`
}
