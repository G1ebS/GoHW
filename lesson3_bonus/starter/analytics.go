package analytics

type Sale struct {
	Product          string
	Category         string
	Quantity         int
	UnitPriceKopecks int64
	Cancelled        bool
}

type Report struct {
	RevenueKopecks    int64
	UnitsSold         int
	RevenueByCategory map[string]int64
	TopProduct        string
	InvalidSales      int
}

type TopProduct struct {
	Product  string
	Quantity int64
}

func AnalyzeSales(sales []Sale) Report {
	report := Report{
		RevenueByCategory: make(map[string]int64),
	}
	productMap := make(map[string]int64)
	for _, sale := range sales {
		switch {
		case sale.Cancelled:
			continue
		case sale.Quantity <= 0 || sale.UnitPriceKopecks <= 0 ||
			sale.Product == "" || sale.Category == "":
			report.InvalidSales++
		default:
			var amount int64 = int64(sale.Quantity) * sale.UnitPriceKopecks
			report.RevenueKopecks += amount
			report.UnitsSold += sale.Quantity
			report.RevenueByCategory[sale.Category] += amount
			productMap[sale.Product] += int64(sale.Quantity)
		}
	}
	topProduct := TopProduct{}
	for product, quantity := range productMap {
		if quantity > topProduct.Quantity {
			topProduct.Product = product
			topProduct.Quantity = quantity
		}
		if quantity == topProduct.Quantity {
			topProduct.Product = max(topProduct.Product, product)
		}
	}
	report.TopProduct = topProduct.Product
	return report
}
