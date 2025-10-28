<div align="right">

[![English](https://img.shields.io/badge/Language-English-blue?style=flat-square)](SECURITY.md)

</div>

# Güvenlik Politikası

## 🔍 Güvenlik Açıklarını Bildirme

Güvenliği ciddiye alıyorum. Bulduğunuz sorunları sorumlu bir şekilde bildirdiğiniz için teşekkür ederim.

### 🚨 **LÜTFEN YAPMAYIN:**

- ❌ Public GitHub issue açmayın
- ❌ Açık düzeltilmeden önce güvenlik açığını herkese açık şekilde ifşa etmeyin
- ❌ Güvenlik açığını göstermek için gerekli olandan fazlasını istismar etmeyin

### ✅ **LÜTFEN BUNLARI YAPIN:**

**Bana e-posta gönderin:** [ahmetsplaxtr@gmail.com](mailto:ahmetsplaxtr@gmail.com)

**Raporunuza şunları ekleyin:**
- Güvenlik açığının türü
- Adım adım yeniden oluşturma ile tam açıklama
- Potansiyel etki
- Önerilen düzeltme (eğer varsa)
- İletişim bilgileriniz

## 🕒 Yanıt Zaman Çizelgesi

Hızlı yanıt vermeye kararlıyım:

| Zaman Çizelgesi | Eylem |
|-----------------|--------|
| **24-48 saat** | Raporunuzun ilk teyidi |
| **72 saat** | Değerlendirme ve önem sınıflandırması |
| **7 gün** | İlerleme hakkında düzenli güncellemeler |
| **30-90 gün** | Hedef düzeltme dağıtımı (öneme bağlı) |

## 🎯 Önem Seviyeleri

### Kritik (CVSS 9.0-10.0)
- **Yanıt süresi:** Anında (24 saat içinde)
- **Düzeltme hedefi:** 1-7 gün
- **Örnekler:** Uzaktan kod yürütme, kimlik doğrulama atlama, veri ihlali

### Yüksek (CVSS 7.0-8.9)
- **Yanıt süresi:** 24-48 saat
- **Düzeltme hedefi:** 7-14 gün
- **Örnekler:** SQL enjeksiyonu, XSS, ayrıcalık yükseltme

### Orta (CVSS 4.0-6.9)
- **Yanıt süresi:** 48-72 saat
- **Düzeltme hedefi:** 14-30 gün
- **Örnekler:** CSRF, bilgi ifşası

### Düşük (CVSS 0.1-3.9)
- **Yanıt süresi:** 1 hafta
- **Düzeltme hedefi:** 30-90 gün
- **Örnekler:** Küçük bilgi sızıntıları, düşük etkili hatalar

## 🛡️ Güvenlik Önlemleri

Projelerim güvenlik en iyi pratiklerini uygular:

### Uygulama Güvenliği
- ✅ Girdi doğrulama ve temizleme
- ✅ SQL enjeksiyonu önleme (parametreli sorgular)
- ✅ XSS koruması (İçerik Güvenlik Politikası)
- ✅ CSRF token'ları
- ✅ Oran sınırlama
- ✅ Güvenli şifre hash'leme (bcrypt/argon2)

### Kimlik Doğrulama ve Yetkilendirme
- ✅ Güçlü şifre gereksinimleri
- ✅ Güvenli oturum yönetimi
- ✅ Rol tabanlı erişim kontrolü (RBAC)
- ✅ Uygun yerlerde OAuth 2.0 desteği

### Veri Koruma
- ✅ Yatarken şifreleme
- ✅ Aktarım sırasında şifreleme (TLS 1.3)
- ✅ Düzenli yedeklemeler
- ✅ Veri saklama politikaları

### Kod Güvenliği
- ✅ Statik kod analizi (SAST)
- ✅ Bağımlılık güvenlik taraması
- ✅ Kod review süreci
- ✅ Güvenli kodlama standartları
- ✅ Otomatik güvenlik testleri

## 🔐 Desteklenen Sürümler

Aşağıdaki sürümler için güvenlik güncellemeleri sağlıyorum:

| Sürüm | Destek Durumu |
|-------|---------------|
| En son | ✅ Aktif destek |
| Önceki | ⚠️ Sadece kritik düzeltmeler |
| Eski | ❌ Artık desteklenmiyor |

**Öneri:** Her zaman en son kararlı sürümü kullanın.

## 📜 Sorumlu İfşa

Koordineli güvenlik açığı ifşası uyguluyorum:

1. **Rapor alındı** - Raporunuzu teyit ediyorum
2. **Doğrulama** - Güvenlik açığını doğruluyor ve değerlendiriyorum
3. **Düzeltme geliştirme** - Düzeltme geliştirip test ediyorum
4. **Düzeltme dağıtımı** - Düzeltmeyi production'a dağıtıyorum
5. **Public ifşa** - 90 gün sonra veya düzeltildiğinde (hangisi önceyse)
6. **Kredi** - Güvenlik danışmanlıklarımda size kredi veriyorum (eğer isterseniz)

## 🎖️ Şeref Listesi

Bana yardımcı olan güvenlik araştırmacılarını takdir ediyorum:

<!-- Bu bölüm katkıda bulunanlarla doldurulacak -->

*Proje güvenliğini iyileştirmemize yardımcı olan ilk kişi siz olun!*

## 📚 Güvenlik En İyi Uygulamaları

### Kullanıcılar İçin
- ✅ Güçlü, benzersiz şifreler kullanın
- ✅ Mevcut olduğunda 2FA'yı etkinleştirin
- ✅ Yazılımı güncel tutun
- ✅ Kimlik avı girişimlerine karşı dikkatli olun
- ✅ Bağlı uygulamaları düzenli olarak gözden geçirin
- ✅ Şüpheli davranışları bildirin

### Geliştiriciler İçin
- ✅ Güvenli kodlama standartlarını takip edin
- ✅ Kod review'lere katılın
- ✅ Testlerde güvenliği ciddiye alın
- ✅ Bağımlılıkları güncel tutun
- ✅ CI/CD'de güvenlik tarayıcılarını kullanın

## 🚩 Bilinen Güvenlik Açıkları

Bilinen sorunlar hakkında şeffaflık sağlıyorum:

**Mevcut Durum:** Bilinen kritik güvenlik açığı yok ✅

Güncellemeler için [Güvenlik Danışmanlıklarını](https://github.com/splaxtr/.github/security/advisories) kontrol edin.

## 📞 İletişim Bilgileri

- **Güvenlik E-posta:** [ahmetsplaxtr@gmail.com](mailto:ahmetsplaxtr@gmail.com)
- **PGP Anahtarı:** Talep üzerine mevcut
- **Yanıt Süresi:** 24-48 saat

## 🚨 Acil Durum Prosedürleri

### Kritik Bir Güvenlik Açığı Keşfederseniz
1. Hemen e-posta gönderin: [ahmetsplaxtr@gmail.com](mailto:ahmetsplaxtr@gmail.com)
2. Mümkünse PGP ile şifreleyin
3. Public ifşadan kaçının
4. Daha fazla işlem yapmadan önce onay bekleyin

### Aktif İstismar Durumunda
1. Derhal haberdar edileceğim
2. Etkilenen kullanıcılar bilgilendirilecek
3. Düzeltme planı paylaşılacak
4. Önleyici tedbirler güçlendirilecek

---

**Son Güncelleme:** Ocak 2025
**Sürüm:** 1.0

Projelerimi ve kullanıcılarını güvende tutmama yardımcı olduğunuz için teşekkür ederim! 🙏
