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

package reqres

type GenericResponse struct {
	Err string `json:"err,omitempty"`
}

// ContentTypeProblemRequest is returned by the ContentTypeValidatingMiddleware
// if there is an unexpected content type.
// Check this in your endpoints to avoid runtime panic when there is a
// mismatched content type.
//
// 		return func(_ context.Context, request interface{}) (interface{}, error) {
//			gr, ok := request.(middleware.ContentTypeProblemRequest)
//
//			if ok {
//				return transport.IsActiveResponse{
//					Err: gr.Err,
//				}, nil
//			}
//			…
type ContentTypeProblemRequest struct {
	Err string `json:"error"`
}