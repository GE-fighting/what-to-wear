'use client';
import { Footer } from '@/components/Footer';

export default function MainPage() {
  return (
    <div className="bg-background-light dark:bg-background-dark font-display">
      <header className="bg-card-light dark:bg-card-dark shadow-sm sticky top-0 z-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-8">
              <h1 className="text-xl font-bold text-text-light-primary dark:text-text-dark-primary">
                StyleSense
              </h1>
              <nav className="hidden md:flex space-x-8">
                <a
                  className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary"
                  href="/record-style/my-outfit"
                >
                  记录穿搭
                </a>
                <a
                  className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary"
                  href="/main/wardrobe"
                >
                  我的衣橱
                </a>
                <a
                  className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary"
                  href="#"
                >
                  风格灵感
                </a>
                <a
                  className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary"
                  href="#"
                >
                  穿搭分析
                </a>
              </nav>
            </div>
            <div className="flex items-center space-x-4">
              <div className="relative hidden md:block">
                <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary">
                  search
                </span>
                <input
                  className="bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-full py-2 pl-10 pr-4 text-sm w-48 focus:ring-primary focus:border-primary text-text-light-primary dark:text-text-dark-primary"
                  placeholder="搜索衣物"
                  type="text"
                />
              </div>
              <div className="relative">
                <a
                  href="/notifications"
                  className="p-2 rounded-full text-text-light-secondary dark:text-text-dark-secondary hover:bg-gray-100 dark:hover:bg-gray-800 inline-block"
                >
                  <span className="material-icons-outlined">notifications_none</span>
                </a>
              </div>
              <a href="/settings/personal-information">
                <img
                  alt="User avatar"
                  className="h-8 w-8 rounded-full object-cover cursor-pointer hover:opacity-80 transition-opacity"
                  src="https://lh3.googleusercontent.com/aida-public/AB6AXuA_qyQReAQqPnp4kLTe4O7X0yHidcXaj2eUnJyZ1VtAgg8kmw6iegfQgMEs2lGGO6nBxZos29reVGCsOSVzLa_LURh9KLREupnhZs9zsi-1CrBrM6Bbf25eIXlGpazovZNI0Xg8J37PVviC-gd0qT2Uj-SQJkN1ihlAb4-fBjEAKwYzFHMXFyHzL6MYZ1pI67jUsw6c5uJ1qqU3-_RZfPsnAu6JWFwAPVgKBP68lF2jyBKS-XrxKrL-9AEOHcFatvQEO6sOzYQxggKL"
                />
              </a>
            </div>
          </div>
        </div>
      </header>

      <main className="h-[calc(100vh-68px)] snap-y overflow-y-scroll">
        <section className="h-full flex items-center justify-center snap-start relative">
          <div className="text-center">
            <h2 className="text-3xl font-bold mb-2 text-text-light dark:text-text-dark">今日推荐穿搭</h2>
            <p className="text-subtext-light dark:text-subtext-dark mb-12">下滑查看下一套</p>
            <div className="grid grid-cols-3 gap-8 max-w-4xl mx-auto items-start">
              <div className="group cursor-pointer">
                <div className="relative bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-lg aspect-square flex items-center justify-center transition-transform duration-300 group-hover:-translate-y-2">
                  <div className="absolute top-3 left-3 space-x-1">
                    <span className="bg-green-100 text-green-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-green-900 dark:text-green-300">
                      棉质
                    </span>
                    <span className="bg-blue-100 text-blue-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-blue-900 dark:text-blue-300">
                      休闲
                    </span>
                  </div>
                  <img
                    alt="简约白T"
                    className="max-h-full w-auto"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuAm-wn4xivBx-CZgIQSyjeoSZbemeyCD-gT9KkLty3fr4q27_KOBc9SbAIfQWyAnHPPW_PR1Mby660c3Bs6LPPVxYCCy7DByd5ggG2J1-NDjeQ_HhYKdbAQ6tGM0r_seL0Heh5qDqU3dhoXGHNUrR0AYWzaRYpOG_5LS33Cq0fgdacRddmC8joofC9DlzIVjnOtxzUoE2NvxUliSw66uzCKHg9xHrrI6_w2fnsMHsAR6Bs8u73FZbW8ZnPL2YfVHir1arzrjbrcJl_X"
                  />
                </div>
                <p className="mt-4 font-semibold text-lg text-text-light dark:text-text-dark">简约绿T</p>
                <p className="text-sm text-subtext-light dark:text-subtext-dark">上装</p>
              </div>
              <div className="group cursor-pointer">
                <div className="relative bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-lg aspect-square flex items-center justify-center transition-transform duration-300 group-hover:-translate-y-2">
                  <div className="absolute top-3 left-3 space-x-1">
                    <span className="bg-yellow-100 text-yellow-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-yellow-900 dark:text-yellow-300">
                      牛仔
                    </span>
                    <span className="bg-purple-100 text-purple-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-purple-900 dark:text-purple-300">
                      街头
                    </span>
                    <span className="bg-pink-100 text-pink-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-pink-900 dark:text-pink-300">
                      春/秋
                    </span>
                  </div>
                  <img
                    alt="直筒牛仔裤"
                    className="max-h-full w-auto"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuD7Rk9xrMR3zea__-Js_HP2eTulRDWT9MTWOsMiJ-Nv5dV1F1yOd3tJs9NDn4vVcI2q3DVwCHUufSQtEbkeyKMcAZ5UWl8OfBzTPsxEYqUpP-sDYsSPsS5drgqPDc_e5uskfIHHHLNGBCxNOpzFJEpzZEHZn_km0ye-foLhox7LXz2JZ_HGOrWkv2mhR9zvwO7m6ZZ5pqYGu8tIWJJtpN6jv7VBz15L3kL4Qzc_wUzQFCXI3CGPntubP-1LunP3BRvaUcVo3ppcmBpN"
                  />
                </div>
                <p className="mt-4 font-semibold text-lg text-text-light dark:text-text-dark">直筒牛仔裤</p>
                <p className="text-sm text-subtext-light dark:text-subtext-dark">下装</p>
              </div>
              <div className="group cursor-pointer">
                <div className="relative bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-lg aspect-square flex items-center justify-center transition-transform duration-300 group-hover:-translate-y-2">
                  <div className="absolute top-3 left-3 space-x-1">
                    <span className="bg-gray-100 text-gray-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-gray-700 dark:text-gray-300">
                      皮革
                    </span>
                    <span className="bg-indigo-100 text-indigo-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-indigo-900 dark:text-indigo-300">
                      运动
                    </span>
                  </div>
                  <img
                    alt="白色运动鞋"
                    className="max-h-full w-auto"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuCK0LqRK0rbCTkblfo4ZjrAk-M4Q-32u9LjsYV1FbBaJcNvn-W19iJPJ49d7B2g2bFW0vT_g2AvQ_COhRvQ_V44bjWLITEcd4UEnMyJD9GmN_wZL4zrQ6aqUDw1VGHVPWGjN23uwnt1m439umV1G0FvjQ-TGERjbamz1SMmXaYW5FbsAaRBbT7iGalIwmLTBrJNd-ypnx8I2GteRLDXbFztTxAx0lz1Q95aEuu39OWabKZ8Ivc02Re-Kx0D2Sm8AcodajJQfogZREWI"
                  />
                </div>
                <p className="mt-4 font-semibold text-lg text-text-light dark:text-text-dark">白色运动鞋</p>
                <p className="text-sm text-subtext-light dark:text-subtext-dark">鞋履</p>
              </div>
            </div>
            <button className="mt-12 bg-primary text-white dark:text-text-light font-bold py-3 px-8 rounded-full inline-flex items-center space-x-2 hover:bg-opacity-90 transition-colors">
              <span className="material-icons">auto_awesome</span>
              <span>AI 试穿</span>
            </button>
          </div>
          <div className="absolute bottom-8 left-1/2 -translate-x-1/2 flex flex-col items-center space-y-1 opacity-70">
            <span className="text-subtext-light dark:text-subtext-dark text-sm">下滑</span>
            <span className="material-icons animate-bounce text-subtext-light dark:text-subtext-dark">south</span>
          </div>
        </section>

        <section className="h-full flex items-center justify-center snap-start relative bg-background-light dark:bg-background-dark">
          <div className="text-center">
            <h2 className="text-3xl font-bold mb-2 text-text-light dark:text-text-dark">机能风穿搭</h2>
            <p className="text-subtext-light dark:text-subtext-dark mb-12">上滑返回 | 下滑查看更多</p>
            <div className="grid grid-cols-3 gap-8 max-w-4xl mx-auto items-start">
              <div className="group cursor-pointer">
                <div className="relative bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-lg aspect-square flex items-center justify-center transition-transform duration-300 group-hover:-translate-y-2">
                  <div className="absolute top-3 left-3 space-x-1">
                    <span className="bg-gray-100 text-gray-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-gray-700 dark:text-gray-300">
                      尼龙
                    </span>
                    <span className="bg-orange-100 text-orange-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-orange-900 dark:text-orange-300">
                      户外
                    </span>
                  </div>
                  <img
                    alt="机能风冲锋衣"
                    className="max-h-full w-auto"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuDCRYtYnl__mP7R1Ozv1_kRWKJpK7qhn08U0Vq8p9HEwtRuYjTbI9--bYRpzVTfKvYEIEK-VtLNOBXHnenYBsGOPg1XZeU2QwKxCGMIzaEG0K7uYkXmiyRfqFQxw37CarGlztEofugPgD4HEDLMqjE6NQwS0UFXFxXM9Ct8N2QLXNoknpQilQQRzDAgEPoosUcfEvqeppBl216RQI7BvogTr68lw4ZRJBIA5ait4-nmE8og44Ir-v1Be36ykNLEMCg5WgklUk_t13hr"
                  />
                </div>
                <p className="mt-4 font-semibold text-lg text-text-light dark:text-text-dark">机能风冲锋衣</p>
                <p className="text-sm text-subtext-light dark:text-subtext-dark">外套</p>
              </div>
              <div className="group cursor-pointer">
                <div className="relative bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-lg aspect-square flex items-center justify-center transition-transform duration-300 group-hover:-translate-y-2">
                  <div className="absolute top-3 left-3 space-x-1">
                    <span className="bg-gray-100 text-gray-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-gray-700 dark:text-gray-300">
                      工装
                    </span>
                    <span className="bg-purple-100 text-purple-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-purple-900 dark:text-purple-300">
                      多口袋
                    </span>
                  </div>
                  <img
                    alt="工装长裤"
                    className="h-full w-full object-cover rounded-lg"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuAIrsUERHX2gv2EfcSvUJl5c7GP0DTEX2dISdPSI6gwvVeEQcX32W9XCY4eUfS3JUdjMrGia1YH9zikCqF5bAevowWyMnNaaJSVdXCzW7S57X11qCVn1p2bWbqbUcLLZj7b4J-U8P8uN3DhriJyUfuiUPk1V3nUYUTOegPHfW7GHo8M5aIR7BgadJ0KYo1AhtHPZWDaWa38k6pXfCcOZ9PZg_96soxuQ46T6bsXAmr3Ym4vJPu41MXxcvGD0mm8JXsK85JMerI4N-bW"
                  />
                </div>
                <p className="mt-4 font-semibold text-lg text-text-light dark:text-text-dark">工装长裤</p>
                <p className="text-sm text-subtext-light dark:text-subtext-dark">下装</p>
              </div>
              <div className="group cursor-pointer">
                <div className="relative bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-lg aspect-square flex items-center justify-center transition-transform duration-300 group-hover:-translate-y-2">
                  <div className="absolute top-3 left-3 space-x-1">
                    <span className="bg-red-100 text-red-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-red-900 dark:text-red-300">
                      高帮
                    </span>
                    <span className="bg-indigo-100 text-indigo-800 text-xs font-medium px-2 py-1 rounded-full dark:bg-indigo-900 dark:text-indigo-300">
                      户外
                    </span>
                  </div>
                  <img
                    alt="高帮运动鞋"
                    className="h-full w-full object-cover rounded-lg"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuAVFLKqiSNrCv_9egIyNCzPdmc20btd2KoSWhr2jSM8ju_L7BsFhZ1-H_lJSxT-SoLy53P8H57YvyCBscSS7qawFARZidK9MAW8Q0bhpOdD1iVzJBDsh_gBVDj7ml0HCfVEnXEWv8bs7F-3V2EV-RTMyAAO9-QpWQoUAksrWFc-wlbylKsRY6xPeX6olr25z91f_uhcwxs3VEQz6zZdx5rKH1zeCua7oJctHbCTIsMcOlqkXNL82-F5iP32nkkt48SorSRYbj9QHhzI"
                  />
                </div>
                <p className="mt-4 font-semibold text-lg text-text-light dark:text-text-dark">高帮运动鞋</p>
                <p className="text-sm text-subtext-light dark:text-subtext-dark">鞋履</p>
              </div>
            </div>
            <button className="mt-12 bg-primary text-white dark:text-text-light font-bold py-3 px-8 rounded-full inline-flex items-center space-x-2 hover:bg-opacity-90 transition-colors">
              <span className="material-icons">auto_awesome</span>
              <span>AI 试穿</span>
            </button>
          </div>
          <div className="absolute bottom-8 left-1/2 -translate-x-1/2 flex flex-col items-center space-y-1 opacity-70">
            <span className="material-icons text-subtext-light dark:text-subtext-dark">north</span>
            <span className="text-subtext-light dark:text-subtext-dark text-sm">上滑/下滑</span>
            <span className="material-icons animate-bounce text-subtext-light dark:text-subtext-dark">south</span>
          </div>
        </section>

        <section className="h-full flex flex-col justify-center snap-start bg-background-light dark:bg-background-dark py-16">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 w-full">
            <div className="mb-12">
              <div className="flex justify-between items-center mb-6">
                <h3 className="text-2xl font-bold text-text-light dark:text-text-dark">流行穿搭</h3>
                <button className="flex items-center space-x-2 text-sm text-subtext-light dark:text-subtext-dark hover:text-primary dark:hover:text-white">
                  <span>换一批</span>
                  <span className="material-icons text-base">refresh</span>
                </button>
              </div>
              <div className="grid grid-cols-2 md:grid-cols-4 gap-6">
                <div className="group relative rounded-lg overflow-hidden cursor-pointer">
                  <img
                    alt="Street style"
                    className="w-full h-64 object-cover transition-transform duration-300 group-hover:scale-105"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuAAbdBK36QF-qFlyXRxgXrlRhBOqfQ2lVWXphg5Y5-trGLQcWYgcaXKr-yYdupFa-YNflAF4Yj12YxLFbncUSAiiaKwOqyxZQftAxl3fCU8colk8jwZcQpDutcydoLcbchsk5uoSKW0NtiI-UFS_jQ4hf6swaXhWb2OAwyYJgJs48GJ-_GWGr9xJpT_sHv0bixyFDj5LaudqrI62T-7g589oEsyXBP0UknISbH5Zj1UPpPRL3cMbk9-Zt87VemnYXhQDO_9bjK06RNq"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <p className="absolute bottom-4 left-4 text-white font-semibold text-lg">街头潮流</p>
                </div>
                <div className="group relative rounded-lg overflow-hidden cursor-pointer">
                  <img
                    alt="Minimalist style"
                    className="w-full h-64 object-cover transition-transform duration-300 group-hover:scale-105"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuC0dWFu2YUzB-SePZBjNpKmaBhbv1c8nxGOGoMJOLjg-EsfWGHfQ1GhGkRKLrC3strbU3n0xVzUi7Vjmjdw7ohqgMb5ebkhhn0GqAPD4u85RxHHocd6rnqMcb0OLBKiAOmcExGfmkKgj5nsqqInkSkCPZjtFx8Y2QTBsZPyOLmfqg7wryyQw0JCVlrBwqLSTaAaKtEQiSqNH0EUQ-hCZ_a9SxLE0VerEBHTGdKIPWtjIUuezqUE5h-8MXez1d1X66WsVB1zMhVDN1NC"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <p className="absolute bottom-4 left-4 text-white font-semibold text-lg">极简主义</p>
                </div>
                <div className="group relative rounded-lg overflow-hidden cursor-pointer">
                  <img
                    alt="Vintage style"
                    className="w-full h-64 object-cover transition-transform duration-300 group-hover:scale-105"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuBvCCqjJgohtrHaISrvWn-Eg8Uk4sithrPZUYthwQmOCGi8hc0kpYmkUkiA9aKA-DAHIrhYA7cYlmBwSRVRQcxA3XPCkH8iO9tdBWGFTD0yQkxBVPnQZ3mqzgpWPiCE106McDGsDnL7h74tXEOBdFM1_AC4Vy5XhxrF7K_Pmv4QJ-QXRvrt2Dv-53VdSWVBwPHuu1vP8QFgY-0iJ3lXPNhb0M-dg6o78s6dmkml_XMEtLWzGryQZ29zDj2tASgk2rnW2dL2DXMWVeJW"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <p className="absolute bottom-4 left-4 text-white font-semibold text-lg">复古风情</p>
                </div>
                <div className="group relative rounded-lg overflow-hidden cursor-pointer">
                  <img
                    alt="Business casual"
                    className="w-full h-64 object-cover transition-transform duration-300 group-hover:scale-105"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuCngJvePkDG2s_LKQP91CFgIcaRWonZ0RF0KJ7M9M5v0AW3Me4OR2hn1CVi_mB-Rnjc8D6RrYdmEaTwacKE1BEMnLKp7pIJvQh-ydYCvR1B0fkYDRJUf8kUAKqNiyU7vB2rJ_o37R0PI5gDl0K25Op3zUr0_o2TxKss7ij3-uHvkzsrUVWnDZ6PehjQ0IPYPdVlTH-e0SS3dN5WWHFeqOq_lGVjMk6_-f_FdVlhiR437_zeaqfn76FY8bplgFIjbi1zJlPA2tH2zeii"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <p className="absolute bottom-4 left-4 text-white font-semibold text-lg">商务休闲</p>
                </div>
              </div>
            </div>
            <div>
              <div className="flex justify-between items-center mb-6">
                <h3 className="text-2xl font-bold text-text-light dark:text-text-dark">我的收藏</h3>
                <button className="flex items-center space-x-2 text-sm text-subtext-light dark:text-subtext-dark hover:text-primary dark:hover:text-white">
                  <span>换一批</span>
                  <span className="material-icons text-base">refresh</span>
                </button>
              </div>
              <div className="grid grid-cols-2 md:grid-cols-4 gap-6">
                <div className="group relative rounded-lg overflow-hidden cursor-pointer">
                  <img
                    alt="Summer vibe outfit"
                    className="w-full h-64 object-cover transition-transform duration-300 group-hover:scale-105"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuB8eUjYGbpzr4ph7KTofzCmV2psn7HgpX8gZd_OV7qwWZXWqrWgvBs82NcT3kv0zSS5AdGLIuDeTe_4tLFl9Gz1ZOvFmzAqEHCSHqz7SyEw8Wlu4VhoZ1D9Ibet821rpRKVlzq3jVuFsP4_c0kkxvpU5_dwBewxCh8IRK3SHP5L0_V9DEhMShsRmDlp-fO__PmV38O727hKKeJd2TWL9ivH179V60ETsOfStiH3YF0YOoYk__Puy8-1rGAeGBHkTmxjAVIsezDrEpdu"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <p className="absolute bottom-4 left-4 text-white font-semibold text-lg">夏日 vibe</p>
                </div>
                <div className="group relative rounded-lg overflow-hidden cursor-pointer">
                  <img
                    alt="Autumn layers"
                    className="w-full h-64 object-cover transition-transform duration-300 group-hover:scale-105"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuAEkMnvRrWPUWYj-IL9KlQVKQFJC-4OqBGai2v1XBoi7LVGSiKH1V3wim8-nfkQjl-0UuUgwg1vvzchGLhGJG_SmOkI6TIcS8AElu74GQtRhKJ0lqLdWG4UbDMmtuSssl52rCI8Mwz7G56SQjzynQXobQEO6lLtgEoI-7MUrQ17oEYvhNx6NxWuL70OjF8356GbypuJ1WiV47VLfXTulaoDiHA1HRcfjPtkPRTlAeK8FzsBfEeEVYxZDc1q5PQbfRzYGhwiLq0WkyEr"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <p className="absolute bottom-4 left-4 text-white font-semibold text-lg">秋日叠穿</p>
                </div>
                <div className="group relative rounded-lg overflow-hidden cursor-pointer">
                  <img
                    alt="Sporty outfit"
                    className="w-full h-64 object-cover transition-transform duration-300 group-hover:scale-105"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuDbbNr3YMXN6437E2SETmn7-lX1xbolRJ2YVzciGuecB7Ns5P13y_fMvgWIIMhBIxkASi7vhWsIuujVtVDQywbQLyBfggHtujPQ6GRyQl_XEjxKj01RyK96nAdkVISF7qBqJt1UJYnRWro9NjDNoaqLRkV63RcibxlESAHfaPJAp34aDSymobVCTR5nh3eGdcRPiTAIwHl7rJPDyZ6_ndESyemGOqgEOHHnBwfN0D-lZjaRDUv0d9mkQGLnEVv8Gv1EGDdEaHkx19oq"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <p className="absolute bottom-4 left-4 text-white font-semibold text-lg">运动休闲</p>
                </div>
                <div className="group relative rounded-lg overflow-hidden cursor-pointer">
                  <img
                    alt="Formal attire"
                    className="w-full h-64 object-cover transition-transform duration-300 group-hover:scale-105"
                    src="https://lh3.googleusercontent.com/aida-public/AB6AXuDtWun_OldpM311_ce0P7wFuIfMZpw4maU_uH49om3fxrFTS7CHzUwXTyDXUtamkWgXX7-KHxuWSDBwm54ki8y2lYyme5zhML3TB0s2PuwxRRo1mTNAGYu55vjZL4RMbucGnOxp7o2exJGVapeBJBSWDugu3s3-5byJrToLaFb7UBsrAXV6FvvcQFRCfPcVsk3RPZ2kwZXCpURHHIIkbL6zKREdYGxfrfrG4gxhWZbBOdpN1Nq-GChWM5b3YXuuO94qgljub8sAUb8W"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <p className="absolute bottom-4 left-4 text-white font-semibold text-lg">正装时刻</p>
                </div>
              </div>
            </div>
          </div>
          <Footer variant="dark" />
        </section>
      </main>
    </div>
  );
}
