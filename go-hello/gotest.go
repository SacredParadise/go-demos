package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"
)

type Flags uint

var prereqs = map[string][]string {
	"algorithm": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}




func main() {
	//randomChan()
	//
	//cannel := make(chan bool)
	//go worker(cannel)
	//
	//time.Sleep(time.Second)
	//cannel <- true

	//errTest()

	//utf8test()
	//
	//fmt.Println(intsToString([]int{1, 2, 3}))

	//str_conv_test()

	//err := eofTest()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//topSort(prereqs)
	//
	//const day = 24 * time.Hour
	//day.Seconds()

	//sortTest()

	//typeAssertTest()

	//selectTest()

	goidtest()
}


func goidtest() {
	panic("goid")
}

func selectTest() {
	ch := make(chan int)

	go func() {
		for i := 1; i < 50; i++ {
			ch <- i
			if i % 20 == 0 {
				time.Sleep(20 * time.Second)
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println(time.Now())
			return
		case c := <-ch:
			fmt.Println(c)
			fmt.Println(time.Now())
		}
	}

}

func typeAssertTest() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	println(f)
}

type StringSlice []string
func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortTest() {
    names := StringSlice {"LiLei", "HanMeimei"}
    sort.Sort(names)
}

func structTest() {
	type Point struct {
		X, Y float64
	}

	type ColoredPoint struct {
		Point
		Color color.RGBA
	}


}

func panicTest(input string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()

	return err
}

func topSort(m map[string][]string) []string {
	var orders []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if (!seen[item]) {
				seen[item] = true
				visitAll(m[item])
				orders = append(orders, item)
			}
		}
	}

	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)

	fmt.Println(orders)
	return orders
}



func eofTest() error {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read failed:%v", err)
		}

		fmt.Println(r)
	}

	return nil
}

func str_conv_test() {
	x := 123
	fmt.Println(strconv.Itoa(x))
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}

		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	return buf.String()
}

func utf8test() {
	var s string = "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println()
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	fmt.Println()
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println()
	fmt.Printf("%x \n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))

	s2 := "abc"
	b := []byte(s2)
	fmt.Println(b)
	s3 := string(b)
	fmt.Println(s3)
	strings.Contains(s, "world")



}

func randomChan() {
	ch := make ( chan int )
	go func () {
		for
		{
			//当有多个管道均可操作时, select 会随机选择一个管道
			select
			{
			case ch <- 0 :
			case ch <- 1 :
			}
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

func worker(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			//正常工作
			case <- cannel:
				//退出
		}
	}
}

func errTest() {
	err := syscall.Chmod(":invalid path:", 0666)
	if err != nil {
		log.Fatal(err.(syscall.Errno))
	}

	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
}


func CopyFile(dstname, srcname string) (written int64, err error) {
	src, err := os.Open(srcname)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstname)
	if err != nil {
		return
	}

	defer dst.Close()

	return io.Copy(dst, src)
}