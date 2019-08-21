package date_utils_test

import (
	"date-utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("DateUtils", func() {
	Context("#date_utils.GetDaysBetween", func() {
		It("should return 3 between 2019-08-12 and 2019-08-15", func() {
			t1, _ := time.Parse("2006-01-02 15:04:05", "2019-08-12 12:00:00")
			t2, _ := time.Parse("2006-01-02 15:04:05", "2019-08-15 17:00:00")
			Expect(date_utils.GetDaysBetween(t1, t2)).Should(Equal(3))
		})

		It("should return 3 between 2019-08-01 and 2019-08-15", func() {
			t1, _ := time.Parse("2006-01-02 15:04:05", "2019-08-01 12:00:00")
			t2, _ := time.Parse("2006-01-02 15:04:05", "2019-08-15 17:00:00")
			Expect(date_utils.GetDaysBetween(t1, t2)).Should(Equal(14))
		})

		It("should return 3 between 2019-07-25 and 2019-08-15", func() {
			t1, _ := time.Parse("2006-01-02 15:04:05", "2019-07-25 12:00:00")
			t2, _ := time.Parse("2006-01-02 15:04:05", "2019-08-15 17:00:00")
			Expect(date_utils.GetDaysBetween(t1, t2)).Should(Equal(21))
		})

		It("should return 3 between 2018-08-15 and 2019-08-15", func() {
			t1, _ := time.Parse("2006-01-02 15:04:05", "2018-08-15 12:00:00")
			t2, _ := time.Parse("2006-01-02 15:04:05", "2019-08-15 17:00:00")
			Expect(date_utils.GetDaysBetween(t1, t2)).Should(Equal(365))
		})

		It("should return 3 between 2017-08-15 and 2019-08-15", func() {
			t1, _ := time.Parse("2006-01-02 15:04:05", "2017-08-15 12:00:00")
			t2, _ := time.Parse("2006-01-02 15:04:05", "2019-08-15 17:00:00")
			Expect(date_utils.GetDaysBetween(t1, t2)).Should(Equal(730))
		})
	})

	Context("#date_utils.NextDay", func() {
		It("should equal 2019-08-13", func() {
			t1, _ := time.Parse("2006-01-02 15:04:05", "2019-08-12 12:00:00")
			t2 := date_utils.NextDay(t1)
			Expect(t2.Format("2006-01-02")).Should(Equal("2019-08-13"))
		})

		It("should equal 2019-08-01", func() {
			t1, _ := time.Parse("2006-01-02 15:04:05", "2019-07-31 12:00:00")
			t2 := date_utils.NextDay(t1)
			Expect(t2.Format("2006-01-02")).Should(Equal("2019-08-01"))
		})

		It("should equal 2019-01-01", func() {
			t1, _ := time.Parse("2006-01-02 15:04:05", "2018-12-31 12:00:00")
			t2 := date_utils.NextDay(t1)
			Expect(t2.Format("2006-01-02")).Should(Equal("2019-01-01"))
		})
	})

	Context("#date_utils.parseExcelNumberToTime", func() {
		It("Should get 2019-04-09 when parse 43564", func() {
			Expect(date_utils.ParseExcelNumber(43564).Format("2006-01-02")).Should(Equal("2019-04-09"))
		})

		It("Should get 2019-05-21 when parse 43606", func() {
			Expect(date_utils.ParseExcelNumber(43606).Format("2006-01-02")).Should(Equal("2019-05-21"))
		})

		It("Should get 2019-07-30 when parse 43676", func() {
			Expect(date_utils.ParseExcelNumber(43676).Format("2006-01-02")).Should(Equal("2019-07-30"))
		})

		It("Should get 2019-08-05 when parse 43682", func() {
			Expect(date_utils.ParseExcelNumber(43682).Format("2006-01-02")).Should(Equal("2019-08-05"))
		})

		It("Should get 2019-08-11 when parse 43688", func() {
			Expect(date_utils.ParseExcelNumber(43688).Format("2006-01-02")).Should(Equal("2019-08-11"))
		})

		It("Should get 2019-08-19 when parse 43696", func() {
			Expect(date_utils.ParseExcelNumber(43696).Format("2006-01-02")).Should(Equal("2019-08-19"))
		})
	})
})
