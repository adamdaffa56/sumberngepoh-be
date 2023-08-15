package helper

import (
	"web-desa/config"
	"web-desa/model"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes),err
}

func SeederRefresh(cfg config.Config) {

	cfg.Database().Migrator().DropTable(&model.User{})
	cfg.Database().Migrator().DropTable(&model.Desa{})
	cfg.Database().Migrator().DropTable(&model.InfoKegiatan{})
	cfg.Database().Migrator().DropTable(&model.Umkm{})
	cfg.Database().Migrator().DropTable(&model.Wisata{})

	hashAdmin1, _ := HashPassword("Admin1**.")
	hashAdmin2, _ := HashPassword("Admin2**.")
	
	users := []model.User{
		{Username: "Admin1", Password: hashAdmin1},
		{Username: "Admin2", Password: hashAdmin2},
	}

	desa := model.Desa{
		ID: 1,
		TentangDesa: "Desa Sumberngepoh merupakan sebuah desa yang terletak di wilayah Kecamatan Lawang, Kabupaten Malang. Desa Sumberngepoh dapat ditempuh dengan waktu sekitar 10 menit dari kantor kecamatan Lawang. Desa Sumberngepoh memiliki luas wilayah seluas 708,60 Ha dan berada di ketinggian 491,00 meter di atas permukaan laut dengan curah hujan 143,00 mm/thn. Di Desa Sumberngepoh terdapat 3 dusun, yaitu Krajan, Barek dan Gapuk. Di bagian utara, Desa Sumber Ngepoh berbatasan dengan Desa Purwodadi. Di bagian selatan, berbatasan dengan Desa Sidodadi. Di bagian timur, berbatasan dengan Desa Cowek. Di bagian barat, berbatasan dengan Desa Mulyoarjo dan Sumberporong.",
		Visi: "Terbangunnya Tata Kelola Pemerintahan Desa yang baik dan Bersih guna mewujudkan Desa Sumberngepoh yang ADIL, MAKMUR, SEJAHTERA dan BERSIH serta BERMARTABAT.",
		Misi: 	"1.) Menyelenggarakan Pemerintahan Desa yang Bersih, Demokratis dan terbebas dari Korupsi, Kolusi dan Nepotisme seta bentuk - bentuk penelewengan lainnya; 2.) Melanjutkan program yang telah dilaksanakan oleh Pemerintah Desa Sumberngepoh periode lalu sebagaimana tercantum dalam Dokumen RPJMDes Desa Sumberngepoh; 3.) Memberdayakan semua potensi vang ada di Masyarakat, meliputi: Pemberdayaan Sumber Daya Manusia (SDM), Pemberdayaan Sumber Daya Alam (SDA) dan, Pemberdavaan Ekonomi Kerakvatan; 4.) Meningkatkan Mutu Kesejahteraan Masyarakat untuk mencapai taraf Kehidupan yang lebih baik dan Berpendidikan; 5.) Menciptakan kondisi Masyarakat Desa yang AMAN, TERTIB, GUYUB RUKUN dalam Kehidupan bermasyarakat dengan berpegangan pada Prinsip - prinsip, yaitu : Duduk sama rendah berdiri sama tinggi, Ringan sama dijinjing, berat sama dipikul.",
	}

	infoKegiatans := []model.InfoKegiatan{
		{
			Judul: "MMD UB Berikan Edukasi Tentang Pentingnya Mengendalikan Tekanan Darah", 
			Gambar: "https://lyofcfhcwphyoxpfuszq.supabase.co/storage/v1/object/public/api-service-desa/informasi-kegiatan/Penyuluhan%20Hipertensi.jpg", 
			Tanggal: "12 Juli 2023", 
			Deskripsi: "Tim Mahasiswa Membangun Desa Universitas Brawijaya (MMD UB) Kelompok 212 Desa Sumberngepoh, Kecamatan Lawang, Kabupaten Malang melakukan kegiatan dengan tema Peningkatan Peran Masyarakat Dalam Pengendalian Penyakit Kronis Untuk Mencapai Kehidupan Sehat dan Sejahtera."+
			"Berdasarkan hasil pemeriksaan yang dilakukan di 5 RW di desa Sumberngepoh didapatkan masalah kesehatan prioritas yaitu “Hipertensi”. Masalah hipertensi ini menjadi prioritas juga di wilayah Lawang, didapatkan dari hasil wawancara bahwa sebagaian masyarakat tidak menjaga pola hidup yang sehat seperti jarang berolahraga dan pola makan yang sembarangan tanpa memperhatikan apakah makanan ini justru dapat meningkatkan tekanan darah. "+
			"Kamis (06/07/2023) Kegiatan ini dimulai dari koordinasi dengan Tenaga Kesehatan di desa. Setelah melakukan koordinasi dan menyesuaikan terkait materi dan metode yang akan digunakan, Tenaga kesehatan desa memberikan saran “kegiatan tersebut sebaiknya dilakukan bersamaan dengan kegiatan posyandu lansia yang dilakukan di dusun Krajan desa Sumberngepoh”." +
			"Selasa (12/07/2023) Acara diawali dengan pemeriksaan kesehatan yang dilakukan oleh perawat dan kader desa Sumberngepoh, Kemudian masyarakat berkumpul dan mengikuti kegiatan penyuluhan yang dilakukan. Mahasiswa membagikan leaflet pada masyarakat, dilanjutkan dengan penyampaian materi oleh Shinta, Materi yang diberikan berupa pengertian hipertensi, cara mengendalikan hipertensi, cara memeriksa tekanan darah mandiri, dan menu makanan yang dianjurkan untuk penderita hipertensi. Untuk susunan acara penyuluhan ini terdiri dari pembukaan, penyampaian materi, tanya jawab, dan penutup. " +
			"Saat dilakukan penyuluhan dengan media leaflet dan poster pada masyarakat di dusun Krajan cukup aktif, dari 23 responden terdapat 5 pertanyaan yang diberikan pada mahasiswa terkait hipertensi. Berdasarkan hasil pre-post test yang dijawab oleh masyarakat didapatkan perbedaan yang signifikan antara tingkat pengetahuan akhir dengan tingkat pengetahuan awal pada responden. Kuesioner diberikan sebagai bentuk evaluasi terlaksananya kegiatan, dari hasil kuesioner didapatkan 20 orang menjawab kegiatan ini bermanfaat dan penting dilakukan. " +
			"Kami berharap ilmu yang telah diberikan dapat bermanfaat bagi masyarakat dan bisa diaplikasikan di kehidupan sehari-hari sehingga terjadi perubahan pola hidup dan terkendalinya hipertensi di desa Sumberngepoh.",
		},
	}

	umkms := []model.Umkm{
		{Nama: "Pawon Anusapati", Alamat: "Dusun Krajan, Desa Sumberngepoh, Kecamatan Lawang", Kontak: "@pawon_anusapati (Instagram), inside.sumberngepoh (Shopee), @inside.sumberngepoh (Instagram)", Gambar: "https://lyofcfhcwphyoxpfuszq.supabase.co/storage/v1/object/public/api-service-desa/umkm/bebek.png", Deskripsi: "UMKM Pawon Anusapati adalah sebuah IKM atau Industri Kecil Menengah yang mengelola bebek dan ayam ungkep di Desa Sumber Ngepoh yang dibuat di Dusun krajan. Bebek Ungkep pawon anusapati menyediakan menu bebek  dan ayam kampung  dalam bentuk frozen food ungkep instan 1 ekor yang siap goreng atau panggang dan bisa disajikan kurang dari 5 menit.. Nama pawon Anusopati diambil dari kata Pawon yaitu Dapur dan Anusapati diambil dari nama raja di kerajaan Singosari UMKM Pawon Anusapati pertama kali didirikan pada tahun 2018. Pendirian UMKM Pawon Anusapati ini dicetuskan oleh Ibu Neni dibantu oleh suami. Pengerjaannya masih sistem manual dan menggunakan rempah rempah khas indonesia tanpa bahan pengawet. bahan bahan diperoleh dari   Tujuan didirikannya UMKM Pawon Anusapati untuk menciptakan peran usaha mikro, kecil dan menengah, penciptaan lapangan kerja, pemerataan pendapatan serta pertumbuhan ekonomi khusunya di pedesaan  UMKM ini masih bersifat mandiri. "},
		{Nama: "Snack Bu Lis", Alamat: "Dusun Krajan, Desa Sumberngepoh, Kecamatan Lawang.", Kontak: "inside.sumberngepoh (Shopee), @inside.sumberngepoh (Instagram)", Gambar: "https://lyofcfhcwphyoxpfuszq.supabase.co/storage/v1/object/public/api-service-desa/umkm/Snack%20Bu%20Lis.png?t=2023-08-01T03%3A01%3A29.488Z", Deskripsi: "UMKM Snack Bu Lis adalah sebuah IKM atau Industri Kecil Menengah yang mengelola berbagai macam keripik di Dusun Krajan Desa Sumber Ngepoh. Snack Bu Lis mulai dirintis pada tahun 2012 oleh Ibu Suliswati sendiri. Awalnya Snack Bu Lis hanya menyediakan produk olahan keripik talas, seiring berjalannya waktu Ibu Suliswati mulai menyediakan olahan keripik tempe dan keripik pisang.  Produksi Snack Bu Lis dilakukan dirumah. Awalnya, penjualan Snack Bu Lis dimulai dari mulut ke mulut, kemudian Ibu Suliswati mulai menerima pesanan dalam jumlah yang banyak dari tetangga-tetangga. "},
		{Nama: "Ikan Nila Bumbu Sumberngepoh", Alamat: "Dusun Krajan, Desa Sumberngepoh, Kecamatan Lawang.", Kontak: "inside.sumberngepoh (Shopee), @inside.sumberngepoh (Instagram)", Gambar: "https://lyofcfhcwphyoxpfuszq.supabase.co/storage/v1/object/public/api-service-desa/umkm/ikan%20nila.png", Deskripsi: "UMKM Ikan Nila Bumbu Sumberngepoh adalah sebuah IKM atau Industri Kecil Menengah yang menjual produk makanan, yaitu ikan nila bumbu. UMKM Ikan Nila Bumbu Sumberngepoh baru didirikan pada tahun 2023 oleh Bapak Andi. Produksi dilakukan dirumah dengan bumbu olahan sendiri tanpa bahan pengawet. Ikan diperoleh dari  budidaya kolam ikan nila  milik desa yang berada di desa sumber ngepoh. ikan nila bumbu frozen food ini dibuat dengan tujuan agar memudahkan  para  konsumen karena  produk yang diolah tanpa bahan pengawet ikan nila bumbu frozen food bisa bertahan hingga 3 bulan apabila disimpan dalam freezer sehingga lebih praktis, mudah, dan juga higienis."},
		{Nama: "Keripik Barokah", Alamat: "Dusun Krajan, Desa Sumberngepoh, Kecamatan Lawang.", Kontak: "inside.sumberngepoh (Shopee), @inside.sumberngepoh (Instagram)", Gambar: "https://lyofcfhcwphyoxpfuszq.supabase.co/storage/v1/object/public/api-service-desa/umkm/keripik%20barokah.png", Deskripsi: "UMKM Keripik Barokah adalah sebuah IKM atau Industri Kecil Menengah yang mengelola berbagai macam keripik di Desa Sumberngepoh. Bisnis keripik ini mulai dirintis oleh Bapak Sujarwoto dari tahun 2008 , dalam pembuatannya dibantu oleh istri dan anaknya, kemudian  mulai di merk kan pada tahun 2014. Bisnis Keripik Barokah milik Pak Sujarwoto ini merupakan salah satu perintis bisnis keripik pertama di Desa Sumberngepoh. Mulai tahun 2014 bisnis Keripik Barokah milik Bapak Sujarwoto ini cukup sukses dengan penjualan yang cukup luas dan memiliki beberapa pegawai dan sales. Keripik Barokah milik Bapak Sujarwoto ini menyediakan aneka keripik singkong dan pisang yang diolah  tanpa bahan pengawet. instagram : @inside.sumberngepoh"},
	}

	wisatas := []model.Wisata{
		{Nama: "Wisata Air Krabyakan", HargaTiket: 5000, Alamat: "Dusun Krajan, Desa Sumber Ngepoh, Kecamatan Lawang, Kabupaten Malang, Jawa Timur.", Kontak: "-", Gambar: "https://lyofcfhcwphyoxpfuszq.supabase.co/storage/v1/object/public/api-service-desa/wisata/krabyakan.jpg", JamBuka: "09:00", JamTutup: "16:00", Deskripsi: "Wisata Air Krabyakan merupakan objek wisata yang terletak di Desa Sumber Ngepoh, Kecamatan Lawang, Kabupaten Malang, Jawa Timur. Objek wisata ini menawarkan pemandian bernuansa pedesaan dengan sumber mata air jernih yang bersih. Daya tarik utama dari Wisata Air Krabyakan adalah keindahan mata air dan kolam yang berdasar bebatuan alami. Wisata air ini menawarkan pesona alam yang memikat berupa hamparan sawah dan perbukitan. Wisata Air Krabyakan memiliki 3 kolam yakni kolam besar utama dimana anak-anak bisa berenang dan bermain air dan terdapat ikan terapi kecil dimana mereka gemar memakan lapisan kulit ari yang mati pada permukaan kaki, kolam air mancur, dan yang terakhir adalah kolam dimana kita bisa bermain sepeda air."},
	}

	cfg.Database().Create(&desa)

	for _, user := range users {
		cfg.Database().Create(&user)
	}

	for _, infoKegiatan := range infoKegiatans {
		cfg.Database().Create(&infoKegiatan)
	}

	for _, umkm := range umkms {
		cfg.Database().Create(&umkm)
	}

	for _, wisata := range wisatas {
		cfg.Database().Create(&wisata)
	}

}