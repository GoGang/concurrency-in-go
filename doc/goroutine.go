func producer(c chan int) {
    for {
        c <- 1
        time.Sleep(100 * time.Nanosecond)
    }
}

func consumer(c chan int) {
    ticker := time.Tick(5 * time.Millisecond)
    result := 0
    for {
        select {
        case val := <-c:
            result = result + val
        case <-ticker:
            println(result)
            return
        }
    }
}
