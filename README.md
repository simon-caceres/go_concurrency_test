# go_concurrency_test
working with concurrency in go


 In order to make a particular function run concurrently, all we have to do is prepend the word "go" to the function call, and it cheerfully runs in the background, as a GoRoutine. Go's built in scheduler takes are of making sure that a given GoRoutine runs when it should, and as efficiently as it can.
