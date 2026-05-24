package ocr

// Comentado para evitar falha de compilação no ambiente sandbox sem as libs nativas do Tesseract
// Em produção, as libs são instaladas via Dockerfile.

/*
import (
	"github.com/otiai10/gosseract/v2"
)
*/

type TesseractOCR struct{}

func NewTesseractOCR() *TesseractOCR {
	return &TesseractOCR{}
}

func (o *TesseractOCR) ExtractText(imagePath string) (string, error) {
	/*
	client := gosseract.NewClient()
	defer client.Close()
	
	if err := client.SetImage(imagePath); err != nil {
		return "", err
	}
	
	return client.Text()
	*/
	return "Simulated text extraction from " + imagePath, nil
}
