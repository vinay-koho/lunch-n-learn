package main

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"
)

var text = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi congue massa mauris, a laoreet mi gravida at. Nam pharetra mattis massa quis lobortis. Sed rutrum arcu eu odio rhoncus ultricies. Pellentesque in neque ut lacus finibus molestie. Etiam sodales interdum est, quis commodo magna auctor quis. Donec eu cursus urna, eu lacinia velit. Sed aliquam ipsum sit amet velit sollicitudin blandit. Curabitur eu risus vel nulla blandit rhoncus. Donec feugiat vestibulum posuere. Quisque eget sem volutpat, suscipit neque a, condimentum arcu. Mauris sed justo vitae neque varius finibus. Phasellus at elit eu quam feugiat aliquam eu id lectus. Nunc consequat risus imperdiet gravida aliquet. Etiam interdum malesuada sem, non auctor eros rutrum vel. Etiam at pellentesque neque, et varius sapien. Vivamus pharetra mauris justo, at dignissim erat fermentum sed.
Donec sit amet ipsum a risus mattis facilisis vitae in ipsum. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nunc ac interdum tortor. Cras pharetra eget erat eget tempus. Nullam posuere rutrum lorem interdum dictum. Nullam dolor sem, sagittis sit amet est feugiat, mollis molestie purus. Quisque facilisis efficitur dui, sit amet mattis nunc egestas non. Nullam sem mi, egestas vehicula dui vitae, interdum condimentum ex.
Sed aliquam dolor vitae nunc rutrum, ut venenatis justo sodales. Suspendisse et turpis eu mauris lacinia finibus in ac justo. Aliquam pellentesque sem diam, id cursus sapien rhoncus ac. Cras vitae tempus ipsum. Maecenas pulvinar, ipsum bibendum commodo fermentum, velit mi laoreet ante, at pharetra ante tortor eget nisi. Curabitur tempor, risus in congue tincidunt, orci ligula ullamcorper enim, feugiat placerat massa sem in lectus. Duis nulla nulla, ullamcorper in condimentum pellentesque, venenatis eget lectus. Morbi placerat neque vel nulla convallis volutpat. Interdum et malesuada fames ac ante ipsum primis in faucibus. Duis urna nibh, eleifend id aliquet at, mollis a urna. Aenean bibendum iaculis lacus et maximus. Nam dictum et magna in ultricies. Nunc accumsan sem tristique tellus dictum malesuada at ut erat. Nunc suscipit posuere est sed laoreet.
Vivamus auctor luctus varius. Maecenas turpis urna, fermentum ac suscipit eget, cursus ac eros. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Praesent vitae libero eget dui consectetur lacinia. Duis maximus eget eros at scelerisque. Morbi commodo rutrum commodo. Vestibulum mollis leo vitae lacus sollicitudin, vel convallis mi consequat.
Vestibulum convallis euismod quam, eget faucibus nisl scelerisque at. Nunc viverra enim fermentum sem maximus, vitae dictum libero interdum. Praesent semper dolor ac purus faucibus fermentum. Nam dapibus elementum commodo. Aliquam fermentum nunc velit, sit amet molestie erat convallis vitae. Donec eu mollis ex, eget dignissim augue. Nullam eget luctus lectus, nec varius est. Donec ornare ex nibh, a molestie nisi aliquam eget.
`

func wordGenerator(lineChannel <-chan string, wordChannel chan<- string, done <-chan bool) {
	defer wg.Done()
	re := regexp.MustCompile(`[^A-z]`)
	for {
		select {
		case line, ok := <-lineChannel:
			if !ok {
				return
			}
			for _, word := range strings.Split(line, " ") {
				wordChannel <- re.ReplaceAllString(word, "")
			}
		case <-done:
			fmt.Println("wordGenerator quiting")
			return
		}
	}
}

func wordCounter(wordChannel <-chan string, done <-chan bool) {
	defer owg.Done()
	for {
		select {
		case word, ok := <-wordChannel:
			if !ok {
				return
			}
			time.Sleep(10 * time.Millisecond)
			wordCountMap[word] = wordCountMap[word] + 1
		case <-done:
			fmt.Println("wordCounter quiting")
			return
		}
	}
}

func monitorInterrupt(done chan<- bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	fmt.Println("System interrupt quiting")
	close(done)
}

var wordCountMap = map[string]int{}
var wg, owg = sync.WaitGroup{}, sync.WaitGroup{}

func main() {
	workerCount := 3
	lineChannel, wordChannel := make(chan string, workerCount), make(chan string, workerCount*1000)
	done := make(chan bool)

	go monitorInterrupt(done)
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go wordGenerator(lineChannel, wordChannel, done)
	}

	owg.Add(1)
	go wordCounter(wordChannel, done)

	for _, line := range strings.Split(text, "\n") {
		lineChannel <- line
	}

	fmt.Println("loaded lines")

	close(lineChannel)
	wg.Wait()

	fmt.Println("populated words")

	close(wordChannel)
	owg.Wait()

	fmt.Println(wordCountMap)
}
