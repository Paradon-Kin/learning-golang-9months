# learning-golang-9months
My 9-month journey from Zero to Back-end Developer
## 📅 Learning Log (บันทึกการเรียนรู้)

| Day | Date | Topic (หัวข้อที่เรียน) | Project (ผลงานที่ทำ) | Status | Key Takeaways (สิ่งที่ได้เรียนรู้) |
| :---: | :---: | :--- | :--- | :---: | :--- |
| **01** | 2026-01-02 | Introduction to Go, Variables, Input/Output | 1. BMI Calculator<br>2. Ideal Weight Challenge | ✅ | เรียนรู้ `var`, `fmt.Scan`, `func main` และการแปลงหน่วย cm/m |
| **02** | - | Data Types & Type Conversion | Currency Converter (ร้านแลกเงิน) | 🔜 | (รอเติม...) |

# Currency Converter / เครื่องแปลงสกุลเงิน

A command-line currency converter built with Go that supports bidirectional conversion between THB and major foreign currencies.

เครื่องแปลงสกุลเงินแบบ command-line ที่สร้างด้วย Go รองรับการแปลงสองทิศทางระหว่างบาทและเงินต่างประเทศหลัก

## Features / ฟีเจอร์

- ✅ Bidirectional conversion (THB ↔ Foreign Currency)
- ✅ Support for USD, JPY, EUR
- ✅ Percentage-based transaction fee (1%)
- ✅ Input validation
- ✅ Clear bilingual interface

## How to Run / วิธีรัน
```bash
go run main.go
```

## Technologies / เทคโนโลジี

- Go 1.23+
- Standard library only (no external dependencies)

## Project Structure / โครงสร้างโปรเจกต์
```
.
├── main.go          # Main program
└── README.md        # This file
```

## Exchange Rates / อัตราแลกเปลี่ยน

As of January 2025:
- 1 THB = 0.029 USD
- 1 THB = 4.32 JPY
- 1 THB = 0.027 EUR

## Author / ผู้เขียน

Part of my 9-month journey to become a Go Backend Developer for G-Able.

เป็นส่วนหนึ่งของการเดินทาง 9 เดือนเพื่อเป็น Go Backend Developer สำหรับ G-Able

## License

MIT