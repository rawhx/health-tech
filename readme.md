
# Health Tech

Aplikasi sederhana untuk memantau mood

## Tools

| Kategori | Tool / Library |
|-----------|----------------|
| Bahasa Pemrograman | Go (Golang) | 
| Web Framework | Gin |
| ORM (Database Layer) | GORM | 
| Database | MySQL | 
| Sanitize | bluemonday |
| Testing Tools | Postman |

## Alasan 

Golang dipilih karena : 

- Performa yang tinggi
- Sederhana
- Stabil

Penggunaan database menggunakan MySQL karena,

- Stabil
- Konsistensi Data
- Ekosistem Luas

## Rencana

- Integrasi dengan sistem authentikasi
- Optimasi Performa
##  Installasi

Clone project

```bash
  git clone https://github.com/rawhx/health_tech
  cd health_tech
```

Setup Environment

```bash
    cp .env.example .env
```

or

```bash
    copy .env.example .env
```    

Install Dependensi

```bash
    go mod tidy
```

Jalankan 

```bash
    go run cmd/main.go
```
## API Endpoint

### Header 
| Key | Values |
|---------|------|
|x-api-key| {apiKey} |

### List Endpoint

-  **Upload Mood**

     Endpoint : `/api/v1/moods`

    Method : `POST`
    
    Body :

    ```json
        {
            "user_id": "f9382d21-ecb5-4b0a-b8c5-fdf1cb4747e6",
            "mood_score": 3,
            "date": "2025-10-27T07:00:00Z",
            "mood_label": "Tenang",
            "notes": "<h1>console.log(1)</h1>"
        }
    ```

    Response :

    ```json
        {
            "status_code": 201,
            "message": "berhasil menambahkan mood"
        }
    ```

- Get Mood

    Endpoint : `/api/v1/moods/:id`

    Method : `GET`

    Parameter :

    | Query | Value |
    | ----- |----- |
    | page | 1 |
    | limit | 10 |

    Response :

    ```json
        {
            "status_code": 201,
            "message": "berhasil menambahkan mood"
        }
    ```

- Get Summary

    Endpoint : `/api/v1/moods/summary/:id`
    
    Method : `GET`

    Param :

    | Query | Value |
    | ----- |----- |
    | periode | week / month |

    Response :

    ```json
        {
            "status_code": 200,
            "message": "berhasil mendapatkan mood summary",
            "data": {
                "data_user": {
                    "user_id": "f9382d21-ecb5-4b0a-b8c5-fdf1cb4747e6",
                    "nama": "abil",
                    "email": "example@gmail.com"
                },
                "rata_rata": 3,
                "totol_data": 62980
            }
        }
    ```

