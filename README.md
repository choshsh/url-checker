# url-checker

Golang으로  http get 방식으로 URL 체크 기능을 만들어봤습니다.

## 설치

[https://github.com/choshsh/url-checker/releases/tag/Release](https://github.com/choshsh/url-checker/releases/tag/Release)

## 사용법

### 파라미터 설정

| Key | Data Type | Description | Example |
| --- | --- | --- | --- |
| url | String | Url to check | -url https://example.com |
| count | Int | (Optional) Number of runs | -count 5 |

### 실행

```bash
url-checker-<OS>-<ARCH> -url <URL> -count <COUNT>
```

### 예시

- MacOS, Linux
    
    ![https://user-images.githubusercontent.com/40452325/150383167-222ce716-5f33-417b-9bbd-13c2acdd477c.png](https://user-images.githubusercontent.com/40452325/150383167-222ce716-5f33-417b-9bbd-13c2acdd477c.png)
    

- Winodws
    
    ![https://user-images.githubusercontent.com/40452325/150383275-96756b40-5c3f-4a1f-9b0e-2fd199c201c9.png](https://user-images.githubusercontent.com/40452325/150383275-96756b40-5c3f-4a1f-9b0e-2fd199c201c9.png)