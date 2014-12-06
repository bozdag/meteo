Tanım
=====

Terminal uygulamalarına küçük bir katkı diyebiliriz. ```cal``` ya da ```alpine``` gibi programları
kullananlar oldukça sevebilirler.

Derleme
=======

Yazılım, Go dilinde geliştirildiği için öncelikle Go dilini indirip kurmalısınız.

Kurulum
=======

İki adet bağımlılığımız var. Bunları kurmak için GOPATH değişkeninizin ayarlı olduğundan emin olun.
Ardından kodlarda yer alan import ifadesinde yer alan ```forecast``` ve ```gcfg``` kütüphanelerini 
indirip kurunuz (```go get``` ile).

En sonunda ```go install``` ile eğer daha önce GOPATH/bin dizinini PATH değişkenine eklediyseniz çalıştırın.


Kullanım
========

```
$ meteo -p <IL PLAKA KODU>
```

Örneğin

```
$ meteo -p 06

Ankara - Parçalı bulutlu bir gece 
-----------------------------------------------------
Sıcaklık(°C) | Orantısal Nem(Y) | Rüzgar Hızı (m/sn)
-----------------------------------------------------
        8.61 |            90.00 |               0.96
```

 Teşekkürler
 ===========

 Hava durumuna ait veriler Forecast.io tarafından sağlanmaktadır.

 API KEY'in Ortak Kullanımı
 ==========================

 API KEY'in ortak kullanılması konusunda benim açımdan bir sorun yok. Sadece günlük bir limitimiz var.
 Bunu bilerek adaletsizce doldurmak son derece *basit* bir davranış olacağından, yapmayın derim :)

 Kendi API Key'inizi edinmek ve kod içerisinde değiştirmek ise süper olur. Tavsiye ederim.