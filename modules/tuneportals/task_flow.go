package tuneportals

// I like this task flow structure, but it doesn't work well with concurrency
// Will need to expand on this...

// The task will pretty much go in this order
// LOAD_CART and LOAD_SHIPPING can be skipped if user provides shipping id
func Start(task *Task) {
	for {
		switch task.Stage {
		case INIT:
			Init(task)
		case PRELOAD_CHECKOUT:
			PreloadCheckout(task)
		case CREATE_SESSION:
			CreateSession(task)
		case LOAD_PRODUCT:
			LoadProduct(task)
		case ADD_TO_CART:
			AddToCart(task)
		case LOAD_CART:
			LoadCart(task)
		case LOAD_SHIPPING:
			LoadShipping(task)
		case TOKENIZE_PAYMENT:
			TokenizePayment(task)
		case SUBMIT_ORDER:
			SubmitOrder(task)
		case KILL:
			return
		}
	}
}
