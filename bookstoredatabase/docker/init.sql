-- สร้างตารางชื่อ books
CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       author VARCHAR(255),
                       isbn VARCHAR(50),
                       year INTEGER,
                       price DECIMAL(10,2),
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- สร้างฟังก์ชันสำหรับอัปเดต updated_at โดยอัตโนมัติ
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

-- สร้างตัวเรียก (Trigger) เพื่อเรียกใช้ฟังก์ชัน update_modified_column
CREATE TRIGGER update_books_modtime
    BEFORE UPDATE ON books
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();

-- สร้างตัวชี้โดยใช้ชื่อหนังสือ เพื่อเพิ่มประสิทธิภาพการค้นหา
CREATE INDEX idx_books_title ON books(title);

-- เพิ่มข้อมูลตัวอย่าง โดยกำหนดลำดับเป็นเลขซีเรียล
INSERT INTO books (title, author, isbn, year, price) VALUES
                                                            ('Fundamental of Deep Learning in Practice', 'Nuttachot Promrit and Sajjaporn Waijanya', '978-1234567890', 2023, 599.00),
                                                            ('Practical DevOps and Cloud Engineering', 'Nuttachot Promrit', '978-0987654321', 2024, 500.00),
                                                            ('Mastering Golang for E-commerce Back End Development', 'Nuttachot Promrit', '978-1111222233', 2023, 450.00);