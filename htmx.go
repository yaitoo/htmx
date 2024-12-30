package htmx

const (

	// HxBoosted indicates that the request is via an element using hx-boost
	HxBoosted = "HX-Boosted"

	// HxCurrentUrl the current URL of the browser
	HxCurrentUrl = "HX-Current-Url"

	// HxHistoryRestoreRequest “true” if the request is for history restoration after a miss in the local history cache
	HxHistoryRestoreRequest = "HX-History-Restore-Request"

	// HxPrompt the user response to an hx-prompt
	HxPrompt = "HX-Prompt"

	// HxRequest always “true” if it is htmx request
	HxRequest = "HX-Request"
	// HxTarget the id of the target element if it exists
	HxTarget = "HX-Target"

	// HxTriggerName the name of the triggered element if it exists
	HxTriggerName = "HX-Trigger-Name"

	// HxTrigger the id of the triggered element if it exists
	HxTrigger = "HX-Trigger"

	// HxLocation allows you to do a client-side redirect that does not do a full page reload
	HxLocation = "HX-Location"

	// HxPushUrl pushes a new url into the history stack
	HxPushUrl = "HX-Push-Url"

	// HxRedirect can be used to do a client-side redirect to a new location
	HxRedirect = "HX-Redirect"

	// HxRefresh if set to “true” the client-side will do a full refresh of the page
	HxRefresh = "HX-Refresh"

	// HxReplaceUrl replaces the current URL in the location bar
	HxReplaceUrl = "HX-Replace-Url"

	// HxReswap allows you to specify how the response will be swapped. See hx-swap for possible values
	HxReswap = "HX-Reswap"
	// HxRetarget a CSS selector that updates the target of the content update to a different element on the page
	HxRetarget = "HX-Retarget"
	// HxReselect a CSS selector that allows you to choose which part of the response is used to be swapped in. Overrides an existing hx-select on the triggering element
	HxReselect = "HX-Reselect"

	// HxTriggerAfterSettle allows you to trigger client-side events after the settle step
	HxTriggerAfterSettle = "HX-Trigger-After-Settle"
	// HxTriggerAfterSwap allows you to trigger client-side events after the swap step
	HxTriggerAfterSwap = "HX-Trigger-After-Swap"
)

type HtmxHeader[T any] map[string]T