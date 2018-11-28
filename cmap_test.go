package cmap_test

import (
	"github.com/kchristidis/cmap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cmap", func() {
	var (
		cm  *cmap.Container
		err error
	)

	Describe("initialization", func() {
		BeforeEach(func() {
			err = nil
		})

		It("error when initialized incorrectly", func() {
			cm, err = cmap.New(0)
			Expect(err).NotTo(BeNil())
			Expect(cm).To(BeNil())
		})

		It("return a non-nil cmap otherwise", func() {
			Expect(err).To(BeNil())
			cm, err = cmap.New(1)
			Expect(err).To(BeNil())
			Expect(cm).NotTo(BeNil())
		})
	})

	Describe("operation", func() {
		BeforeEach(func() {
			cm, _ = cmap.New(2)
		})

		It("put/get a key/value pair", func() {
			cm.Put("fooKey", "fooVal")
			val, ok := cm.Get("fooKey")
			Expect(ok).To(BeTrue())
			Expect(val.(string)).To(Equal("fooVal"))
		})

		It("delete a key/value pair", func() {
			cm.Put("fooKey", "fooVal")
			ok := cm.Delete("fooKey")
			Expect(ok).To(BeTrue())
			val, ok := cm.Get("fooKey")
			Expect(ok).To(BeFalse())
			Expect(val).To(BeNil())
		})

		It("attempt to get a non-existing key", func() {
			val, ok := cm.Get("fooKey")
			Expect(ok).To(BeFalse())
			Expect(val).To(BeNil())
		})

		It("attempt to delete a non-existing key", func() {
			ok := cm.Delete("fooKey")
			Expect(ok).To(BeFalse())
		})

		It("store different types", func() {
			cm.Put("fooKey", 10)
			cm.Put(20, "barValue")
			val, ok := cm.Get("fooKey")
			Expect(ok).To(BeTrue())
			Expect(val).To(Equal(10))
			val, ok = cm.Get(20)
			Expect(ok).To(BeTrue())
			Expect(val).To(Equal("barValue"))
		})

		Context("at capacity", func() {
			BeforeEach(func() {
				cm.Put("fooKey", "fooVal1")
				cm.Put("barKey", "barVal")
			})

			It("replace an existing key", func() {
				cm.Put("fooKey", "fooVal2")
				val, ok := cm.Get("fooKey")
				Expect(ok).To(BeTrue())
				Expect(val.(string)).To(Equal("fooVal2"))
			})

			It("evict the LRU key", func() {
				cm.Get("fooKey")
				cm.Put("bazKey", "bazVal")
				val, ok := cm.Get("barKey")
				Expect(ok).To(BeFalse())
				Expect(val).To(BeNil())
			})
		})
	})
})
