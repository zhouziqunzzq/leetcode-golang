package mycode

import "math/rand"

type Codec struct {
	code2Url   map[string]string
	url2Code   map[string]string
	codeBucket []rune
}

func Constructor535() Codec {
	k := 0
	buf := make([]rune, 26*2+10)
	for i := 0; i < 26; i++ {
		buf[k] = rune('A' + i)
		k++
	}
	for i := 0; i < 26; i++ {
		buf[k] = rune('a' + i)
		k++
	}
	for i := 0; i < 10; i++ {
		buf[k] = rune('0' + i)
		k++
	}

	return Codec{
		code2Url:   make(map[string]string),
		url2Code:   make(map[string]string),
		codeBucket: buf,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	const codeLen = 6
	codeBuf := make([]rune, codeLen)
	for {
		if code, ok := this.url2Code[longUrl]; !ok {
			for i := range codeBuf {
				codeBuf[i] = this.codeBucket[rand.Intn(len(this.codeBucket))]
			}
			randCode := string(codeBuf)
			if _, ok1 := this.code2Url[randCode]; !ok1 {
				this.code2Url[randCode] = longUrl
				this.url2Code[longUrl] = randCode
			}
		} else {
			return "http://tinyurl.com/" + code
		}
	}
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.code2Url[shortUrl[len(shortUrl)-6:]]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
