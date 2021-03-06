# Golang_Goroutine

Goroutine 

- Goroutine adalah sebuah thread yg dikelola Go-Runtime
- Goroutine sangat ringan (2Kb/Goroutine), dibandingkan dengan thread(1MB/Goroutine)
- Goroutine berjalan secara concurrent , thread berjalan paralel

Cara Kerja Goroutine
- Goroutine dijalankan oleh Go-Scheduler, dimana jumlah threadnya sebanyak CPU
- Goroutine berjalan diatas thread
- Namun kita tidak perlu management thread secara manual karena semua diatur oleh Go-Scheduler

Cara Kerja Go-Scheduler
- Dalam Go-Scheduler ada beberapa terminologi
    1. G : Goroutine
    2. M : Thread(Machine)
    3. P : Processor


Channel
    - Channel adalah tempat komunikasi secara synchronous yg dilakukan Goroutine
    - Dalam Channel ada pengirim dan penerima yg biasanya berupa Goroutine yg berbeda
    - Saat pengiriman data ke dalam Channel, Goroutine akan blocking sampai ada yg menerima data tersebut
    - Mekanisme Channel hampir sama dengan async/await pada bahasa pemrograman lain
    - Atau dapat disimpulkan Channel adalah alat komunikasi data synchronous(blocking)

Karakteristik Channel
    - Secara default channel hanya bisa menampung satu tipe data, jika ingin menambah data harus menunggu data tersebut ada yg mengambil
    - Channel hanya bisa menerima 1 tipe data
    - Channel bisa diambil oleh beberapa Goroutine
    - Channel harus di close jika selesai digunakan agar tidak menyebabkan memory leak

Channel sebagai Parameter
    - Dalam praktek pembuatan design app kita akan mengirim channel ke function lain via Parameter
    - Dalam Golang defaultnya parameter adalah 'Pass by Value' artinya value akan di duplikasi lalu dikirim
        ke function parameter, sehingga jika kita ingin mengirim data asli maka harus menggunakan pointer
        agar 'Pass by Reference'
    - Dalam channel kita tidak perlu mengakses pointer Parameter

Buffered Channel

    - Buffered Channel Adalah buffer yg bisa digunakan untuk menampung data Antrian di Channel
    - Kapasitasnya Bebas sesuai kebutuhan channel 
