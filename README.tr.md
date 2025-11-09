<div align="center">
  <h1>splaxtr/.github</h1>
  <p>Tüm projelerde paylaşılan workflow'lar, issue/PR şablonları ve topluluk dosyaları burada tutulur.</p>
</div>

<p align="center">
  <a href="./README.md"><img src="https://img.shields.io/badge/Language-EN-blue?style=for-the-badge" alt="English"></a>
  <a href="./WORKFLOWS_GUIDE.md"><img src="https://img.shields.io/badge/Dok%C3%BCman-Workflows%20Guide-6f42c1?style=for-the-badge&logo=githubactions" alt="Workflows Guide"></a>
  <a href="./SETUP_GUIDE.md"><img src="https://img.shields.io/badge/Dok%C3%BCman-Kurulum%20K%C4%B1lavuzu-0ea5e9?style=for-the-badge" alt="Kurulum Kılavuzu"></a>
  <a href="./FOLDER_STRUCTURE.md"><img src="https://img.shields.io/badge/Dok%C3%BCman-Klas%C3%B6r%20Yap%C4%B1s%C4%B1-16a34a?style=for-the-badge" alt="Klasör Yapısı"></a>
  <a href="./CONTRIBUTING.md"><img src="https://img.shields.io/badge/Dok%C3%BCman-Katk%C4%B1-f97316?style=for-the-badge" alt="Katkı"></a>
  <a href="./SECURITY.md"><img src="https://img.shields.io/badge/Dok%C3%BCman-G%C3%BCvenlik-ef4444?style=for-the-badge" alt="Güvenlik Politikası"></a>
  <a href="https://github.com/splaxtr/.github/actions/workflows/validate-reusables.yml"><img src="https://github.com/splaxtr/.github/actions/workflows/validate-reusables.yml/badge.svg" alt="Doğrulama İş Akışı"></a>
  <a href="https://github.com/splaxtr/.github/actions/workflows/self-lint.yml"><img src="https://github.com/splaxtr/.github/actions/workflows/self-lint.yml/badge.svg" alt="Lint Denetimi"></a>
  <a href="https://github.com/splaxtr/.github/actions/workflows/reusable-security.yml"><img src="https://github.com/splaxtr/.github/actions/workflows/reusable-security.yml/badge.svg" alt="Güvenlik Denetimi"></a>
</p>

## [🇬🇧 English](README.md) · 🇹🇷 Türkçe

### 📚 İçindekiler
1. [Genel Bakış](#genel-bakış)
2. [Depo İçeriği](#depo-içeriği)
3. [Yeniden Kullanılabilir Workflowlar](#yeniden-kullanılabilir-workflowlar)
4. [Dokümantasyon](#dokümantasyon)
5. [Hızlı Başlangıç](#hızlı-başlangıç)
6. [Doğrulama ve Örnek Projeler](#doğrulama-ve-örnek-projeler)
7. [Destek ve İletişim](#destek-ve-iletişim)

---

## Genel Bakış
Bu depo, splaxtr organizasyonu içindeki tüm paylaşılan GitHub varlıklarını barındırır. Workflow'lar, issue/PR şablonları, politikalar ve yardımcı scriptler tek yerde tutulur, böylece alt projeler aynı ayarları kopyalamaya ihtiyaç duymadan tutarlı kalır. Yapı ayrıca yeniden kullanılabilir workflow'ların otomatik olarak doğrulanması için hafif örnek uygulamalar (Node, Python, Go, Flutter, Docker) içerir.

## Depo İçeriği
| Kategori | Sağlananlar |
| --- | --- |
| 🔄 Workflow'lar | Node, Bun, Python, Go, Flutter, Docker build, deploy, metrik, güvenlik taraması ve toplu doğrulama için yeniden kullanılabilir CI/CD pipeline'ları. |
| 📋 Şablonlar | GitHub Issue formları (bug, feature, question) ile sadeleştirilmiş PR şablonu ve CODEOWNERS uyumlu reviewer uyarıları. |
| 📜 Politikalar | CONTRIBUTING, SECURITY, SUPPORT, CODE_OF_CONDUCT, FUNDING ve dil bazlı varyasyonlar ile tüm depolar için uyumluluk. |
| 🏗️ Örnek projeler | `frontend`, `backend`, `mobile`, `bun-app`, `bots/*` ve `docker` klasörleri CI doğrulaması için fixture görevi görür. |
| 🧰 Araçlar | Dependabot, etiket tanımları, release drafter konfigürasyonu, workflow ayarları ve bootstrap otomasyonu için scriptler. |

## Yeniden Kullanılabilir Workflowlar
| Workflow | Açıklama |
| --- | --- |
| [reusable-node-ci](./workflows/reusable-node-ci.yml) | npm, yarn veya pnpm projeleri için lint/test/build matrisleri. |
| [reusable-bun-ci](./workflows/reusable-bun-ci.yml) | Bun projelerinde lint, coverage, build ve isteğe bağlı E2E aşaması. |
| [reusable-python-ci](./workflows/reusable-python-ci.yml) | Pip kurulumları, Ruff/Black kontrolleri, pytest coverage ve artifact yüklemeleri. |
| [reusable-go-ci](./workflows/reusable-go-ci.yml) | GolangCI-Lint, race destekli testler, binary build ve isteğe bağlı Docker imajı. |
| [reusable-flutter-ci](./workflows/reusable-flutter-ci.yml) | Flutter format/analyze/test ve Android/Web build matrisleri. |
| [reusable-docker-build](./workflows/reusable-docker-build.yml) | GHCR push destekli, cache'li Buildx süreçleri ve metadata etiketleme. |
| [reusable-deploy-production](./workflows/reusable-deploy-production.yml) | Backend, frontend, mobil, bot ve veritabanı için çok servisli deploy orkestrasyonu. |
| [reusable-metrics-ci](./workflows/reusable-metrics-ci.yml) | LOC/coverage/dosya sayımı + SVG özet artefaktları. |
| [reusable-monorepo-ci](./workflows/reusable-monorepo-ci.yml) | Monorepo içinde her teknoloji yığını için paralel CI çalıştırır. |
| [reusable-security](./workflows/reusable-security.yml) | npm audit + Trivy dosya sistemi taraması + CodeQL analizi. |
| [ci-lint](./workflows/ci-lint.yml) | Tüm yeniden kullanılabilir parçaları zincirleyen bileşik doğrulama workflow'u. |
| [validate-reusables](./workflows/validate-reusables.yml) | Örnek projelerle tüm workflow'ları çalıştıran depo içi doğrulama işi. |

Detaylı girişler ve sorun giderme notları için **[WORKFLOWS_GUIDE.md](./WORKFLOWS_GUIDE.md)** dosyasına bakın.

## Dokümantasyon
- [WORKFLOWS_GUIDE.md](./WORKFLOWS_GUIDE.md) – Her workflow'un giriş/çıkışları ve kullanım örnekleri.
- [FOLDER_STRUCTURE.md](./FOLDER_STRUCTURE.md) – Şablonlar ve doğrulama varlıkları dahil tüm dosyaların açıklamalı klasör ağacı.
- [SETUP_GUIDE.md](./SETUP_GUIDE.md) – Organizasyon seviyesinde `.github` reposu oluşturma/güncelleme, iletişim bilgilerini özelleştirme ve workflow bağlama adımları.
- [CONTRIBUTING.md](./CONTRIBUTING.md) – Branch modeli, commit konvansiyonları, PR beklentileri ve etiket kullanımı.
- [SECURITY.md](./SECURITY.md) – Açıklama politikası ve `ahmetsplaxtr@gmail.com` iletişim adresi.
- Ek kaynaklar: [labels-README.md](./labels-README.md), [SUPPORT.md](./SUPPORT.md), [SUPPORT.tr.md](./SUPPORT.tr.md).

## Yönetişim & Otomasyon
- [CODEOWNERS](./CODEOWNERS) – İnceleme yönlendirmelerini zorunlu kılar ve sorumlulukları belirtir.
- [dependabot.yml](./dependabot.yml) – CI, frontend ve backend ekosistemlerini kapsayan haftalık bağımlılık güncellemelerini yönetir.
- [release-drafter.yml](./release-drafter.yml) – Semantik sürüm notlarını ve sürüm numaralarını otomatik oluşturur.
- [labels.yml](./labels.yml) – Otomasyonla uyumlu tip/kapsam/öncelik/semver etiket taksonomisini tanımlar.

## Yönetişim ve Otomasyon
- [CODEOWNERS](./CODEOWNERS) – Workflow, şablon, doküman ve örnek projeler için reviewer sahiplikleri.
- [dependabot.yml](./dependabot.yml) – CI, frontend ve backend ekosistemleri için haftalık bağımlılık güncellemeleri.
- [release-drafter.yml](./release-drafter.yml) – Feature/Fix/Docs/Security kategorileriyle semantik sürüm notları.
- [labels.yml](./labels.yml) – Organizasyon genelinde tip, kapsam, öncelik, durum ve semver etiketleri.

## Hızlı Başlangıç
1. **Workflow'u referans alın:** projenize aşağıdaki gibi yeniden kullanılabilir workflow ekleyin.
   ```yaml
   name: CI
   on: [push, pull_request]
   jobs:
     build:
       uses: splaxtr/.github/.github/workflows/reusable-node-ci.yml@main
       with:
         node_version: '20.x'
         working_directory: '.'
   ```
2. **Şablonları kullanın:** depo public ise GitHub yerel şablon yoksa otomatik olarak bu issue/PR şablonlarına başvurur.
3. **Kılavuzları takip edin:** SETUP_GUIDE'de anlatıldığı gibi iletişim adreslerini, politikaları ve dokümanları projelerinize uyarlayın.

## Doğrulama ve Örnek Projeler
- `validate-reusables.yml` workflow'u push/workflow_dispatch tetiklerinde çalışır ve örnek klasörleri (`frontend`, `backend`, `mobile`, `bun-app`, `bots/node`, `bots/python`, `docker`) kullanarak her yeniden kullanılabilir workflow'u test eder.
- Yerelde denemek isterseniz aynı workflow'ları `act` ile koşturabilir veya fork üzerinde `workflow_dispatch` ile tetikleyebilirsiniz.

## Workflow Doğrulaması
Bu depoda `self-lint.yml` ve `validate-reusables.yml` pipeline'ları bulunur; her değişiklikte workflow'ların sözdizimi,
önbellek kullanımı ve çıktı uyumluluğu otomatik olarak doğrulanır.

## Lint Denetimi
- `self-lint.yml`, yamllint ve actionlint ile tüm workflow'ları her PR'da tarar.

## Güvenlik Denetimi
- `reusable-security.yml`, npm audit, Trivy ve CodeQL taramalarını tek bir ortak yapı altında çalıştırır.

## Destek ve İletişim
- Sorular: Discussion açın veya "Question" issue formunu doldurun.
- Hata/özellik: İlgili issue formunu detaylı doldurun.
- Güvenlik: [SECURITY.md](./SECURITY.md) içinde anlatıldığı gibi **ahmetsplaxtr@gmail.com** adresine özel olarak iletin.
- Genel destek: [SUPPORT.md](./SUPPORT.md) dokümanındaki e-posta ve SLA bilgilerini kullanın.

Katkılarınız sayesinde platform tutarlı kalıyor, teşekkürler! 🚀
