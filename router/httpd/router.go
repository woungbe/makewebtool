package httpd

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouters(e *echo.Echo) {
	e.GET("/", html)                     // 일반화면
}

func html(c echo.Context) error {
	return c.Render(http.StatusOK, "html", nil)
}

// func uploadView(c echo.Context) error {
// 	return c.Render(http.StatusOK, "upload", map[string]interface{}{
// 		"url": "/excel/csv/uploadsave",
// 	})
// }

func errorPage(c echo.Context) error {
	return c.Render(http.StatusBadRequest, "error", nil)
}


// //「/save」용 핸들러
// func saveHandler(w http.ResponseWriter, r *http.Request) {
// 	//MultipartReader를 이용해서 받은 파일을 읽는다
// 	reader, err := r.MultipartReader()
 
// 	//에러가 발생하면 던진다
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
 
// 	//for로 복수 파일이 있는 경우에 모든 파일이 끝날때까지 읽는다
// 	for {
// 		part, err := reader.NextPart()
// 		if err == io.EOF {
// 			break
// 		}
 
// 		//파일 명이 없는 경우는 skip
// 		if part.FileName() == "" {
// 			continue
// 		}
 
// 		//uploadedfile 디렉토리에 받았던 파일 명으로 파일을 만든다
// 		uploadedFile, err := os.Create("./upload/" + part.FileName())
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			uploadedFile.Close()
// 			redirectToErrorPage(w,r)
// 			return
// 		}
 
// 		//만든 파일에 읽었던 파일 내용을 모두 복사
// 		_, err = io.Copy(uploadedFile, part)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			uploadedFile.Close()
// 			redirectToErrorPage(w,r)
// 			return
// 		}
// 	}
// 	//upload 페이지에 리다이렉트
// 	http.Redirect(w,r,"/upload",http.StatusFound)
//  }
 
//  //「/upload」용 핸들러
//  func uploadHandler(w http.ResponseWriter, r *http.Request) {
// 	var templatefile = template.Must(template.ParseFiles("./upload/"))
// 	templatefile.Execute(w, "upload.html")
//  }
 
//  //「/errorPage」용 핸들러
//  func errorPageHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w,"%s","<p>Internal Server Error</p>")
//  }
 
//  //error가 발생했을 때에 에러 페이지로 이동한다
//  func redirectToErrorPage(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w,r,"/errorPage",http.StatusFound)
//  }
 
