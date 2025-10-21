

### عنوان
سِینار (Sinar) — سامانه‌ی یکپارچه مدیریت و سفارش غذای دانشگاهی بر بستر Go و Gin

### معرفی کلی
- **هدف**: ارائه‌ی سرویسی برای مدیریت کاربران (دانشجویان)، دانشگاه‌ها، رستوران‌های هر دانشگاه، منوی غذا و خرید/مصرف غذا با احراز هویت OTP.
- **فناوری‌ها**: Go 1.24+، Gin، PostgreSQL، Redis، Swagger.
- **ویژگی‌ها**:
  - **احراز هویت با OTP** (ارسال و تأیید کد)
  - **مدیریت کاربر** بر اساس شماره دانشجویی
  - **مدیریت دانشگاه** و بازیابی اطلاعات
  - **مدیریت رستوران‌ها** به تفکیک دانشگاه
  - **مدیریت غذا** و فهرست غذاها
  - **خرید غذای کاربر** با تاریخ انقضا و امکان «مصرف/باطل‌شدن»

### معماری سامانه
- **الگوی لایه‌ای (Clean/Hexagonal-Style)**:
  - `internal/domain`: مدل‌های دامنه (User, University, Restaurant, Food, UserFood)
  - `internal/dto`: آبجکت‌های انتقال داده برای پاسخ‌های ترکیبی (مثل `UserWithUniversity`, `dto.UserFood`)
  - `internal/usecase`: منطق کسب‌وکار و قراردادها (interfaceهای `Store` برای لایه‌ی زیرساخت)
  - `internal/interface/postgres`: ریپازیتوری‌های پایگاه‌داده PostgreSQL (CRUD/Query)
  - `internal/interface/redis`: ذخیره‌سازی OTP با TTL
  - `internal/interface/server`: کنترلرهای HTTP (Gin Handlers) و مسیرها
  - `internal/config`: بارگذاری پیکربندی Redis و SMS از متغیرهای محیطی
  - `pkg/sms`: درگاه ارسال پیامک OTP (قابل جایگزینی)
  - `docs`: Swagger برای مستندسازی API
  - `main.go`: سیم‌کشی وابستگی‌ها و راه‌اندازی سرور و Swagger

- **جریان وابستگی‌ها**:
  - Handler → Usecase → Store Interface → Repository (Postgres/Redis)
  - مدل‌های دامنه در `domain` مستقل از زیرساخت نگهداری می‌شوند.
  - DTOها برای پاسخ‌های غنی (Join چند موجودیت) استفاده می‌گردند.

- **راه‌اندازی سرویس‌ها** در `main.go`:
  - بارگذاری کانفیگ (`config.New`)، ایجاد `redis` برای OTP و `sms` برای ارسال.
  - ساخت Usecaseها با تزریق ریپازیتوری‌ها.
  - رجیستر کردن روت‌ها و Swagger.

### موجودیت‌ها و روابط (EER)
- **User**
  - فیلدها: `ID`, `FirstName`, `LastName`, `Phone`, `ProfilePic`, `StudentNum`, `Sex`, `UniversityID`
- **University**
  - فیلدها: `ID`, `Name`, `Location`, `Logo`
- **Restaurant**
  - فیلدها: `ID`, `UniversityID`, `Name`, `Sex`, `Color`
- **Food**
  - فیلدها: `ID`, `Name`
- **UserFood**
  - فیلدها: `ID`, `UserID`, `FoodID`, `RestaurantID`, `Price`, `SinarPrice`, `Code`, `CreatedAt`, `ExpiresAt`

روابط و کاردینالیتی:
- University 1 ──< User
  - هر کاربر به یک دانشگاه تعلق دارد؛ هر دانشگاه چندین کاربر دارد.
- University 1 ──< Restaurant
  - هر رستوران متعلق به یک دانشگاه است؛ هر دانشگاه چندین رستوران دارد.
- User 1 ──< UserFood
  - هر خرید غذا توسط یک کاربر انجام می‌شود؛ هر کاربر چندین خرید دارد.
- Food 1 ──< UserFood
  - هر خرید به یک «غذا» اشاره می‌کند؛ هر غذا در خریدهای متعدد استفاده می‌شود.
- Restaurant 1 ──< UserFood
  - هر خرید به یک رستوران اشاره می‌کند؛ هر رستوران خریدهای متعدد دارد.

نکات EER:
- کلیدهای خارجی: `users.university_id → universities.id`، `restaurants.university_id → universities.id`، `user_foods.user_id → users.id`، `user_foods.food_id → foods.id`، `user_foods.restaurant_id → restaurants.id`
- ویژگی‌های زمانی: `user_foods.expires_at` برای مدیریت اعتبار/مصرف.
- کد یکتا برای خرید: `user_foods.code` (برای رهگیری/مصرف).

### گردش‌های کلیدی
- **OTP**:
  - ایجاد کد با طول و TTL مشخص، ذخیره در Redis و ارسال از طریق `pkg/sms`.
  - تأیید: خواندن از Redis و انطباق با ورودی، سپس حذف.
- **دریافت اطلاعات کاربر با دانشگاه**:
  - Join `users` و `universities` و ارائه DTO `UserWithUniversity`.
- **لیست رستوران‌های دانشگاه**:
  - فیلتر بر اساس `university_id`.
- **لیست غذا**:
  - دریافت ساده از جدول `foods`.
- **خرید غذا (UserFood)**:
  - حالت 1 (بدنه تکی): دریافت `expiration_hours` و محاسبه `expires_at`.
  - حالت 2 (آرایه): دریافت `expires_at` و تبدیل به `expiration_hours` تقریبی.
  - درج رکورد و بازگردانی `id` و `created_at`.
- **علامت‌گذاری به‌عنوان مصرف‌شده**:
  - اگر `expires_at` گذشته باشد، یعنی مصرف‌شده/باطل شده؛ در غیر این صورت با عقب‌بردن `expires_at`، آن را مصرف‌شده می‌کنیم.

### APIها (خلاصه)
- OTP:
  - POST `/otp/create` ایجاد و ارسال کد
  - POST `/otp/verify` تأیید کد
- User:
  - GET `/user/{student_number}` اطلاعات کاربر + دانشگاه
- University:
  - GET `/university/{id}` دریافت دانشگاه
- Restaurant:
  - GET `/restaurant/{university_id}` فهرست رستوران‌های یک دانشگاه
- Food:
  - GET `/food/` فهرست غذاها
- UserFood:
  - GET `/userfood/active` خریدهای فعال (منقضی نشده)
  - GET `/userfood/{id}` جزئیات یک خرید
  - POST `/userfood/` ایجاد خرید (تکی یا آرایه)
  - POST `/userfood/{id}/use` مصرف/باطل‌کردن خرید

### لایه Usecase (نمونه قراردادها)
- **User**: `GetByStudentNumber(number string) (*dto.UserWithUniversity, error)`
- **Restaurant**: `GetAll(uniID int) ([]*domain.Restaurant, error)`
- **Food**: `GetAllNames() ([]*domain.Food, error)`
- **UserFood**: `Purchase(...), GetActive(), GetByID(), MarkAsUsed()`
- **OtpService**: `RequestOTP(phone string), VerifyOTP(userID, otp string)`

### داده‌ها و SQL
- اسکریپت‌های پایگاه‌داده در `Database/` موجود است.
- اتصال پایگاه‌داده فعلاً در `internal/interface/postgres/postgres.go` به‌صورت رشته اتصال ثابت تنظیم شده است (برای محیط‌های واقعی پیشنهاد می‌شود از متغیرهای محیطی/Secrets استفاده شود).

### پیکربندی
- `internal/config/load.go` از متغیرهای محیطی می‌خواند:
  - `REDIS_ADDR`, `REDIS_PASS`, `REDIS_DB`
  - `SMS_APIKEY`, `SMS_SENDER`
- پیشنهاد: اضافه‌کردن متغیرهای محیطی برای اتصال PostgreSQL بجای hardcode.

### استقرار و اجرا
- پیش‌نیازها: Go، PostgreSQL، Redis
- گام‌ها:
  - `go mod download`
  - اجرای اسکریپت‌های دیتابیس (`Database/tables.sql`)
  - تنظیم متغیرهای محیطی Redis/SMS
  - اجرای سرویس: `go run main.go`
  - Swagger UI: `http://localhost:8080/swagger/index.html`

### لاگینگ و مانیتورینگ
- لاگ‌های ساده از طریق `log` در OTP؛ می‌توان با `pkg/logger` یکپارچه‌سازی کامل‌تری انجام داد.
- پیشنهاد: اضافه‌کردن متریک‌ها (Prometheus) و رهگیری درخواست‌ها.

### امنیت
- احراز هویت مبتنی بر OTP (شماره تلفن). در Swagger تعریف `BearerAuth` وجود دارد اما در هندلرها اعمال نشده؛ در صورت نیاز می‌توان JWT را یکپارچه کرد.
- اعتبارسنجی پارامترها در Usecaseها و Handlers انجام می‌شود (چک خالی/منفی بودن و ...).

### محدودیت‌ها و پیشنهادهای بهبود
- **Postgres DSN**: خارج‌سازی از کد و استفاده از ENV/Secret.
- **Transactions/Isolation**: افزودن تراکنش در عملیات حساس خرید.
- **Indexes**: ایندکس روی کلیدهای خارجی و ستون‌های پرمصرف مانند `student_num`, `expires_at`.
- **Idempotency**: جلوگیری از خرید تکراری با همان `code`.
- **Observability**: اضافه‌کردن ساختار لاگ استاندارد، TraceID در پاسخ‌ها.
- **Rate Limiting**: محدودسازی درخواست‌های OTP برای هر شماره.
- **Testing**: تکمیل تست‌های واحد و همگرایی برای Usecaseها و Repos.

### EER (تشریح دقیق)
- کلیدهای اصلی: `id` در همه‌ی جداول اصلی.
- کلیدهای خارجی:
  - `users.university_id` → دانشگاه کاربر
  - `restaurants.university_id` → دانشگاه رستوران
  - `user_foods.user_id` → مالک خرید
  - `user_foods.food_id` → نوع غذا
  - `user_foods.restaurant_id` → محل خرید/ارائه
- ویژگی‌های مهم:
  - `user_foods.code`: شناسه رهگیری خرید (پیشنهاد: یکتا)
  - `user_foods.expires_at`: انقضا/مصرف
- کاردینالیتی:
  - 1:N بین University و User/Restaurant
  - 1:N بین User و UserFood
  - 1:N بین Food و UserFood
  - 1:N بین Restaurant و UserFood
- اختیاری/اجباری:
  - `users.university_id` اجباری (به فرض)
  - `user_foods.*_id` اجباری
  - برخی فیلدهای نمایشی مثل `profile_pic` می‌توانند اختیاری باشند.

### استفاده نمونه از مدل‌ها (ارجاع کد)
```1:15:/run/media/danny/8E96D53A96D5238D/work/sinar/internal/domain/models.go
type User struct {
    ID           int    `json:"id"`
    FirstName    string `json:"first_name"`
    LastName     string `json:"last_name"`
    Phone        string `json:"phone"`
    ProfilePic   string `json:"profile_pic"`
    StudentNum   string `json:"student_num"`
    Sex          bool   `json:"sex"`
    UniversityID int    `json:"university_id"`
}
```

```23:46:/run/media/danny/8E96D53A96D5238D/work/sinar/internal/domain/models.go
type UserFood struct {
    ID           int       `json:"id"`
    UserID       int       `json:"user_id"`
    FoodID       int       `json:"food_id"`
    RestaurantID int       `json:"Restaurant_id"`
    Price        int       `json:"price"`
    SinarPrice   int       `json:"sinar_price"`
    Code         string    `json:"code"`
    CreatedAt    time.Time `json:"created_at"`
    ExpiresAt    time.Time `json:"expires_at"`
}
```

### نتیجه‌گیری
سِینار یک سرویس ماژولار و لایه‌ای برای مدیریت فرایندهای سفارش غذای دانشگاهی است که با جداسازی مسئولیت‌ها در `usecase`، `interface` و `domain`، توسعه‌پذیری و نگهداشت را ساده کرده است. زیرساخت Redis برای OTP و Postgres برای داده‌های تراکنشی، و Swagger برای مستندسازی، مسیر رشد و استقرار را هموار می‌کند.

- افزودن پیکربندی پویا برای Postgres، بهبود امنیت (JWT+RBAC)، تراکنش‌ها و Observability، گام‌های بعدی پیشنهاد می‌شوند.

- تغییرات/اثر: مستندات جامع فارسی شامل عنوان، معماری، EER، APIها، جریان‌ها و پیشنهادات بهبود ارائه شد.
