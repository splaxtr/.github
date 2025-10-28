<div align="right">

[![English](https://img.shields.io/badge/Language-English-blue?style=flat-square)](CONTRIBUTING.md)
[![Türkçe](https://img.shields.io/badge/Dil-Türkçe-red?style=flat-square)](CONTRIBUTING.tr.md)

</div>

# Katkıda Bulunma

Öncelikle, projelerime katkıda bulunmayı düşündüğünüz için teşekkür ederiz! 🎉

Katkılarınız bu projeleri herkes için daha iyi hale getiriyor.

## 📋 İçindekiler

- [Davranış Kuralları](#davranış-kuralları)
- [Nasıl Katkıda Bulunabilirim?](#nasıl-katkıda-bulunabilirim)
- [Geliştirme Kurulumu](#geliştirme-kurulumu)
- [Pull Request Süreci](#pull-request-süreci)
- [Kodlama Standartları](#kodlama-standartları)
- [Commit Kılavuzları](#commit-kılavuzları)
- [Test Kılavuzları](#test-kılavuzları)

---

## 📜 Davranış Kuralları

Bu proje ve içinde yer alan herkes [Davranış Kurallarımız](CODE_OF_CONDUCT.md) ile yönetilir. Katılarak bu kurallara uymayı kabul etmiş olursunuz.

---

## 🤝 Nasıl Katkıda Bulunabilirim?

### Hata Bildirme

Hata raporu oluşturmadan önce, tekrarları önlemek için mevcut issue'ları kontrol edin.

**İyi bir hata raporu nasıl yazılır:**
- [Hata raporu şablonunu](.github/ISSUE_TEMPLATE/bug_report.yml) kullanın
- Net bir başlık ve açıklama sağlayın
- Yeniden oluşturma adımlarını ekleyin
- Uygunsa ekran görüntüsü/log ekleyin
- Ortamınızı belirtin (İşletim Sistemi, tarayıcı, versiyon)

### Özellik Önerme

Özellik önerilerini seviyoruz!

**İyi bir özellik isteği nasıl yazılır:**
- [Özellik isteği şablonunu](.github/ISSUE_TEMPLATE/feature_request.yml) kullanın
- Çözmeye çalıştığınız problemi açıklayın
- Önerdiğiniz çözümü tanımlayın
- Alternatifleri düşünün
- Mevcut kullanıcılar üzerindeki etkiyi değerlendirin

### Kod Katkısı

**İlk katkılar için iyi issue'lar:**
[`good first issue`](https://github.com/search?q=user%3Asplaxtr+label%3A%22good+first+issue%22+state%3Aopen&type=Issues) etiketi ile işaretlenmiş issue'ları arayın - yeni başlayanlar için harika!

**Katkıda bulunma adımları:**
1. Repository'yi fork'layın
2. Bir feature branch oluşturun (`git checkout -b feature/harika-ozellik`)
3. Değişikliklerinizi yapın
4. Test yazın/güncelleyin
5. Testleri çalıştırın (`npm test` veya eşdeğeri)
6. Değişikliklerinizi commit'leyin ([commit kılavuzları](#commit-kılavuzları)'nı takip edin)
7. Fork'unuza push'layın (`git push origin feature/harika-ozellik`)
8. Bir Pull Request açın

---

## 🛠️ Geliştirme Kurulumu

### Ön Gereksinimler

Projeye bağlı olarak:
- **Node.js**: 18+ (JavaScript/TypeScript projeleri için)
- **Python**: 3.11+ (Python projeleri için)
- **Go**: 1.21+ (Go projeleri için)
- **Flutter**: 3.13+ (Flutter projeleri için)
- **Docker**: 20+ (opsiyonel, konteynerleştirilmiş geliştirme için)
- **Git**: 2.30+

### Başlangıç

```bash
# Repository'yi klonlayın
git clone https://github.com/splaxtr/[proje-adi].git
cd [proje-adi]

# Bağımlılıkları yükleyin
npm install  # veya: pip install -r requirements.txt, go mod download, flutter pub get

# Testleri çalıştırın
npm test     # veya: pytest, go test ./..., flutter test

# Geliştirmeyi başlatın
npm run dev  # veya projeye özel komut
```

---

## 🔄 Pull Request Süreci

### Göndermeden Önce

1. **Dokümantasyonu güncelleyin** - README, API dokümanları, vb.
2. **Test yazın** - %80+ kod kapsamını koruyun
3. **Linter'ları çalıştırın** - `npm run lint` veya eşdeğeri
4. **Tüm testleri çalıştırın** - Her şeyin geçtiğinden emin olun
5. **CHANGELOG'u güncelleyin** - "Yayınlanmamış" altına giriş ekleyin
6. **Kendi kendinize review yapın** - Önce kendi kodunuzu gözden geçirin

### PR Gereksinimleri

- [ ] Net başlık (conventional commits'e uygun)
- [ ] Açıklama ne ve neden'i açıklıyor
- [ ] Tüm testler geçiyor
- [ ] Kod kapsamı korundu/iyileştirildi
- [ ] Dokümantasyon güncellendi
- [ ] Merge conflict yok
- [ ] En az 1 reviewer tarafından onaylandı

### PR Başlık Kuralı

```
<tip>(<kapsam>): <konu>

Örnekler:
feat(auth): OAuth2 desteği ekle
fix(api): WebSocket'te bellek sızıntısını çöz
docs(readme): kurulum talimatlarını güncelle
refactor(core): performansı iyileştir
test(utils): yardımcılar için birim testleri ekle
```

**Tipler:**
- `feat` - Yeni özellik
- `fix` - Hata düzeltmesi
- `docs` - Sadece dokümantasyon
- `style` - Formatlama, noktalı virgül eksikliği, vb.
- `refactor` - Kod yeniden yapılandırması
- `perf` - Performans iyileştirmesi
- `test` - Test ekleme/güncelleme
- `chore` - Bakım görevleri

---

## 💻 Kodlama Standartları

### Genel İlkeler

- **Basit tutun** - Akıllılık yerine açıklığı tercih edin
- **DRY** - Kendinizi Tekrar Etmeyin
- **YAGNI** - Buna İhtiyacınız Olmayacak
- **Test yazın** - Kodunuzu iyice test edin
- **Dokümante edin** - Sadece "ne"yi değil, "neden"i açıklayın

### JavaScript/TypeScript

**Stil Rehberi:** [Airbnb JavaScript Style Guide](https://github.com/airbnb/javascript)'ı takip edin

```typescript
// ✅ İyi
interface User {
  id: string;
  username: string;
  email: string;
}

async function fetchUser(userId: string): Promise<User> {
  const response = await api.get(`/users/${userId}`);
  return response.data;
}

// ❌ Kötü
function getUser(id) {
  return api.get('/users/' + id).then(r => r.data);
}
```

**Kurallar:**
- TypeScript kullanın (`any` tipi yok)
- `.then()` yerine `async/await` kullanın
- `let` yerine `const`, asla `var` kullanmayın
- Callback'ler için ok fonksiyonları kullanın
- ESLint ve Prettier çalıştırın

### Python

**Stil Rehberi:** [PEP 8](https://peps.python.org/pep-0008/)'i takip edin

```python
# ✅ İyi
def fetch_user(user_id: str) -> User:
    """Kullanıcıyı ID'ye göre getirir.

    Args:
        user_id: Kullanıcının ID'si

    Returns:
        User: Kullanıcı nesnesi

    Raises:
        UserNotFound: Kullanıcı bulunamazsa
    """
    response = api.get(f'/users/{user_id}')
    return User.from_dict(response.json())

# ❌ Kötü
def getuser(id):
    return User.from_dict(api.get('/users/'+id).json())
```

### Go

**Stil Rehberi:** [Effective Go](https://golang.org/doc/effective_go.html)'yu takip edin

```go
// ✅ İyi
func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
    if err := s.validator.Validate(req); err != nil {
        return nil, fmt.Errorf("doğrulama başarısız: %w", err)
    }
    // ... implementasyon
}

// ❌ Kötü
func createUser(r *CreateUserRequest) *User {
    // Eksik context, hata yönetimi, doğrulama
}
```

---

## 📝 Commit Kılavuzları

### Konvansiyonel Commit'ler

[Conventional Commits](https://www.conventionalcommits.org/)'i takip ediyoruz.

**Format:**
```
<tip>(<kapsam>): <konu>

<gövde>

<alt bilgi>
```

**Örnek:**
```
feat(auth): OAuth2 giriş akışı ekle

Google, GitHub ve Microsoft sağlayıcıları ile
OAuth2 authorization code flow'unu uygula.

- OAuth2 istemci konfigürasyonu ekle
- Token değişimini uygula
- Kullanıcı profili getirmeyi ekle
- Giriş UI'sını güncelle

#123'ü kapatır
```

**Kurallar:**
- Şimdiki zaman kullanın ("özellik ekle" "özellik ekledi" değil)
- Emir kipi kullanın ("imleci taşı" "imleç taşır" değil)
- İlk harfi büyük yapmayın
- Sonunda nokta yok
- Alt bilgide issue/PR'lara referans verin

---

## 🧪 Test Kılavuzları

### Test Kapsamı

- **Minimum**: %80 kod kapsamı
- **Hedef**: Kritik kod için %90+

### Test Piramidi

```
       /\
      /  \     E2E (%10)
     /____\
    /      \   Entegrasyon (%30)
   /________\
  /          \ Birim (%60)
 /____________\
```

### Test Yazma

**JavaScript/TypeScript:**
```typescript
describe('fetchUser', () => {
  it('kullanıcıyı başarıyla getirmeli', async () => {
    // Düzenle
    const userId = '123';
    mockApi.get.mockResolvedValue({ data: mockUser });

    // Hareket
    const user = await fetchUser(userId);

    // Doğrula
    expect(user).toEqual(mockUser);
    expect(mockApi.get).toHaveBeenCalledWith('/users/123');
  });
});
```

**Python:**
```python
def test_fetch_user():
    # Düzenle
    user_id = "123"
    mock_api.get.return_value = MockResponse(mock_user)

    # Hareket
    user = fetch_user(user_id)

    # Doğrula
    assert user == mock_user
    mock_api.get.assert_called_with('/users/123')
```

---

## 🏷️ Issue Etiketleri

Issue'ları organize etmek ve önceliklendirmek için etiketleri kullanıyoruz:

**Tip:**
- `type: bug` - Bir şey çalışmıyor
- `type: feature` - Yeni özellik isteği
- `type: docs` - Dokümantasyon iyileştirmesi
- `type: refactor` - Kod yeniden yapılandırması
- `type: performance` - Performans iyileştirmesi

**Öncelik:**
- `priority: critical` - Acil, çekirdek işlevselliği bozuyor
- `priority: high` - Önemli, yakında yapılmalı
- `priority: medium` - Normal öncelik
- `priority: low` - Olsa iyi olur

**Durum:**
- `status: needs-triage` - İlk review gerekiyor
- `status: needs-info` - Daha fazla bilgi bekleniyor
- `status: blocked` - Başka bir issue tarafından engelleniyor
- `status: in-progress` - Şu anda üzerinde çalışılıyor
- `status: needs-review` - Review için hazır

**Özel:**
- `good first issue` - Yeni gelenler için iyi
- `help wanted` - Topluluk yardımı gerekli
- `breaking change` - Breaking change getiriyor
- `security` - Güvenlikle ilgili issue

---

## 🎯 Geliştirme İş Akışı

### Branch Stratejisi

```
main (production)
  ↑
develop (staging)
  ↑
feature/xxx (feature branch'leri)
```

**Branch İsimlendirme:**
- `feature/oauth-ekle` - Yeni özellikler
- `fix/bellek-sizintisi` - Hata düzeltmeleri
- `docs/api-rehberi` - Dokümantasyon
- `refactor/auth-servisi` - Refactoring
- `hotfix/guvenlik-yamasi` - Acil düzeltmeler

---

## ❓ Sorularınız mı var?

- 📧 İletişim bilgileri için repository'nin README'sine bakın
- 💬 Bir [Tartışma](https://github.com/splaxtr/.github/discussions) açın
- 🐛 Hatalar için [Issues](https://github.com/splaxtr/.github/issues) kullanın

---

## 🙏 Teşekkürler!

Katkılarınız bu projeleri herkes için daha iyi hale getiriyor. Zamanınız ve çabanız için minnettarız! ❤️

**İyi kodlamalar! 🚀**
