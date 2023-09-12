package service

import (
	"math/rand"
	"path/filepath"
	"strings"
	"time"
	"web-desa/model"
	"web-desa/request"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
	"github.com/gin-gonic/gin"
)

type infoKegiatanService struct {
	infoKegiatanRepository model.InfoKegiatanRepository
}

func NewInfoKegiatanService(infoKegiatan model.InfoKegiatanRepository) model.InfoKegiatanService {
	return &infoKegiatanService{infoKegiatanRepository: infoKegiatan}
}

func (s *infoKegiatanService) StoreInfoKegiatan(req *request.InfoKegiatanRequest) (*model.InfoKegiatan, error) {
	infoKegiatan := &model.InfoKegiatan{
		Judul: req.Judul,
		Gambar: req.Gambar,
		Tanggal: req.Tanggal,
		Deskripsi: req.Deskripsi,
	}

	newInfoKegiatan, err := s.infoKegiatanRepository.Create(infoKegiatan)
	if err != nil {
		return nil, err
	}

	return newInfoKegiatan, nil
}

func (s *infoKegiatanService) EditInfoKegiatan(id uint, req *request.InfoKegiatanRequest) (*model.InfoKegiatan, error) {
	newDesa, err := s.infoKegiatanRepository.UpdateByID(&model.InfoKegiatan{
		ID: id,
		Judul: req.Judul,
		Gambar: req.Gambar,
		Tanggal: req.Tanggal,
		Deskripsi: req.Deskripsi,
	})

	if err != nil {
		return nil, err
	}

	return newDesa, nil
}

func (s *infoKegiatanService) GetByID(id uint) (*model.InfoKegiatan, error) {
	infoKegiatan, err := s.infoKegiatanRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return infoKegiatan, err
}

func (s *infoKegiatanService) DestroyInfoKegiatan(id uint) error {
	infoKegiatan, errFind := s.infoKegiatanRepository.FindByID(id)
	if errFind != nil {
		return errFind
	}

	_, err := s.infoKegiatanRepository.Delete(infoKegiatan)
	if err != nil {
		return err
	}

	return nil
}

func (s *infoKegiatanService) FetchInfoKegiatan() ([]*model.InfoKegiatan, error) {
	infoKegiatans, err := s.infoKegiatanRepository.Fetch()
	if err != nil {
		return nil, err
	}

	return infoKegiatans, nil
}

var supClient = supabasestorageuploader.NewSupabaseClient(
	"https://lyofcfhcwphyoxpfuszq.supabase.co",
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Imx5b2ZjZmhjd3BoeW94cGZ1c3pxIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTA4NTgwODcsImV4cCI6MjAwNjQzNDA4N30.4ghuCygd-B4EX2NWq93UatzvEuIAFlozmSifDowO1fc",
	"api-service-desa",
	"informasi-kegiatan",
)

func (h *infoKegiatanService) UploadImage(c *gin.Context) (string, error) {
	file, err := c.FormFile("gambar")
	if err != nil {
		return "", err
	}

	// generate randomString
    randomString := RandomString(5)

	// untuk mendapatkan ekstensi file
    ext := filepath.Ext(file.Filename)

	// menghasilkan nama baru dari penggabungan nama file(tanpa ekstensi) + randomString + ekstensi file
    newFilename := strings.TrimSuffix(file.Filename, ext) + randomString + ext

	// inisialisasi Filename dengan fileName baru
    file.Filename = newFilename

	link, err := supClient.Upload(file)
	if err != nil {
		return "", err
	}
	return link, nil
}

func (h *infoKegiatanService) DeleteImage(c *gin.Context, id uint) error {
	infoKegiatan, errFind := h.GetByID(id)
	if errFind != nil {
		return errFind
	}

	_, err := supClient.DeleteFile(infoKegiatan.Gambar)
	if err != nil {
		return err
	} 

	return nil
}

func RandomString(length int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    
	b := make([]rune, length)
    for i := range b {
        b[i] = letters[randomizer.Intn(len(letters))]
    }
    return string(b)
}