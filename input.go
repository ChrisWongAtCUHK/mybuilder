package main

type A struct {
	a string
	b string
	c string
}

type ABuilder struct {
	a string
	b string
	c string
}

func (builder *ABuilder) a(a string) {
	builder.a = a
}

func (builder *ABuilder) b(b string) {
	builder.b = b
}

func (builder *ABuilder) c(c string) {
	builder.c = c
}

func (builder *ABuilder) build() *ABuilder {
	return &ABuilder{
		a: builder.a,
		b: builder.b,
		c: builder.c,
	}
}

