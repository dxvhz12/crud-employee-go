CREATE TABLE IF NOT EXISTS data_karyawan (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    alamat VARCHAR(255),
    jabatan VARCHAR(50),
    gaji NUMERIC(12,2),
    tanggal_masuk DATE,
    tanggal_keluar DATE,
    foto VARCHAR(50),
    status VARCHAR(20) DEFAULT 'aktif'
);