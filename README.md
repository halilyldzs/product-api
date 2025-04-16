# Product API

Bu proje, Go programlama dili kullanılarak geliştirilmiş RESTful bir ürün yönetim API'sidir. API, Fiber web framework'ü, GORM ORM kütüphanesi ve SQLite veritabanı kullanılarak oluşturulmuştur.

## Özellikler

- Ürün CRUD (Create, Read, Update, Delete) işlemleri
- Swagger UI ile API dokümantasyonu
- DTO (Data Transfer Object) modelleri ile veri validasyonu
- GORM ile veritabanı işlemleri
- CGO gerektirmeyen saf Go SQLite sürücüsü
- Fiber web framework ile hızlı ve verimli HTTP işleme
- Soft delete özelliği
- Otomatik zaman damgaları (created_at, updated_at)

## Başlangıç

### Gereksinimler

- Go 1.22 veya üstü

### Kurulum

1. Projeyi klonlayın:

   ```bash
   git clone https://github.com/kullanici/product-api.git
   cd product-api
   ```

2. Bağımlılıkları yükleyin:

   ```bash
   go mod tidy
   ```

3. Projeyi çalıştırın:

   ```bash
   go run main.go
   ```

4. API sunucusu `http://localhost:8080` adresinde çalışmaya başlayacaktır.

## API Dokümantasyonu

API dokümantasyonuna `http://localhost:8080/swagger/index.html` adresinden erişebilirsiniz.

### Endpoints

| Method | Endpoint             | Açıklama                    |
| ------ | -------------------- | --------------------------- |
| POST   | /api/v1/products     | Yeni bir ürün oluştur       |
| GET    | /api/v1/products     | Tüm ürünleri listele        |
| GET    | /api/v1/products/:id | Belirli bir ürünü getir     |
| PUT    | /api/v1/products/:id | Var olan bir ürünü güncelle |
| DELETE | /api/v1/products/:id | Bir ürünü sil               |

## Proje Yapısı

```
product-api/
├── database/           # Veritabanı yapılandırması ve bağlantısı
├── docs/               # Swagger tarafından oluşturulan API dokümantasyonu
├── models/             # Veri modelleri ve DTO'lar
├── repository/         # Veritabanı işlemleri (CRUD)
├── routes/             # API route tanımları ve handler'lar
├── data/               # SQLite veritabanı dosyası (çalışma zamanında oluşturulur)
├── main.go             # Ana uygulama dosyası
└── README.md           # Bu dosya
```

## Modeller

### Product Model

```go
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       float64        `json:"price" gorm:"not null"`
	Quantity    int            `json:"quantity" gorm:"default:0"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	SKU         string         `json:"sku"`
	Barcode     string         `json:"barcode"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
```

### DTO Modelleri

```go
type ProductCreateDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	SKU         string  `json:"sku"`
	Barcode     string  `json:"barcode"`
}

type ProductUpdateDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	SKU         string  `json:"sku"`
	Barcode     string  `json:"barcode"`
}
```

## Örnek İstekler

### Ürün Oluşturma (POST /api/v1/products)

```json
{
  "name": "Örnek Ürün",
  "description": "Bu bir örnek ürün açıklamasıdır",
  "price": 99.99,
  "quantity": 10,
  "sku": "ORN-12345",
  "barcode": "1234567890123"
}
```

### Ürün Güncelleme (PUT /api/v1/products/1)

```json
{
  "name": "Güncellenmiş Ürün",
  "description": "Bu ürün açıklaması güncellenmiştir",
  "price": 149.99,
  "quantity": 5,
  "sku": "ORN-54321",
  "barcode": "3210987654321"
}
```

## Teknolojiler

- [Go](https://golang.org/) - Ana programlama dili
- [Fiber](https://gofiber.io/) - Hızlı, Express benzeri web framework
- [GORM](https://gorm.io/) - Go için ORM kütüphanesi
- [SQLite](https://www.sqlite.org/) - Gömülü SQL veritabanı
- [Swagger](https://swagger.io/) - API dokümantasyonu
- [github.com/glebarez/sqlite](https://github.com/glebarez/sqlite) - CGO gerektirmeyen saf Go SQLite sürücüsü

## Geliştirme

### Swagger Dokümantasyonunu Güncelleme

API dokümantasyonunu güncellemek için:

```bash
swag init
```

### Ek Özellik Önerileri

- Kullanıcı kimlik doğrulama ve yetkilendirme
- Ürün kategorileri
- Gelişmiş arama ve filtreleme
- Resim yükleme desteği
- Sayfalama
- Loglama
- Birim ve entegrasyon testleri

## Lisans

Bu proje [MIT Lisansı](LICENSE) altında lisanslanmıştır.
