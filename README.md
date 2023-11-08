# Üniversite - Fakülte - Bölüm Listesi JSON + SQL 

Bu GitHub deposu, Yükseköğretim Kurulu (YÖK) üzerinden alınan üniversite, fakülte ve bölüm listesini JSON ve SQL formatlarına dönüştürmek için kullanılan bir araç içerir. Bu aracı kullanarak YÖK verilerini kolayca işleyebilir ve istediğiniz veri formatını oluşturabilirsiniz.

## Kullanım

Aşağıda bu aracı kullanmanın temel adımları verilmiştir:

1. **Veri İndirme**: YÖK tarafından https://istatistik.yok.gov.tr adresinden  Birim İstatistikleri > Genel Bilgiler > Bölümler 'den  sağlanan Excel dosyasını indirin ve proje kodlarının olduğu dizine kaydedin.
2. Kod'da fileName alanlarını ve sheetName alanını düzenleyin
3. Uygulamayı çalıştırın. `go run main.go`
4. **Excel Verisinin JSON ve SQL'e Dönüştürülmesi**:

   - Excel verisini bu aracı kullanarak JSON ve SQL formatlarına dönüştürün.
   - Dönüştürme işlemi için kullanılan komutları ve parametreleri açıklamak için gerekirse kod örneklerini ekleyin.
5. **Sonuçları Kullanma**: Dönüştürülen verileri başka projelerde veya uygulamalarda kullanabilirsiniz. Örneğin, SQL sorgularını bir veritabanına yükleyebilir veya JSON verilerini başka sistemlerle entegre edebilirsiniz.

## Katkıda Bulunma

Bu projeye katkıda bulunmak isterseniz, lütfen aşağıdaki adımları izleyin:

1. Bu repo'yu çatallayın (fork).
2. Yeni bir özellik eklemek, hata düzeltmek veya iyileştirmeler yapmak için yeni bir dal (branch) oluşturun.
3. Değişikliklerinizi yapın ve düzenleyin.
4. Değişikliklerinizi test edin ve gerekirse belgeleri güncelleyin.
5. Değişikliklerinizi bu repo'ya bir çekme isteği (pull request) gönderin.
