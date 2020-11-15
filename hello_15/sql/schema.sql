CREATE DATABASE IF NOT EXISTS finance;

-- 表引擎 Memory 重启后原数据消失

CREATE TABLE finance.brch_qry_dtl (
    acc String,
    tran_date Date,
    amt Decimal(10, 2),
    dr_cr_flag Int,
    rpt_sum String,
    timestamp1 String
) ENGINE = Memory;

--分区建表 表引擎 MergeTree

CREATE TABLE finance.brch_qry_dtl (
    tran_date Date,
    timestamp1 String,
    acc String,
    amt Decimal(10, 2),
    dr_cr_flag Int,
    rpt_sum String
) ENGINE = MergeTree
PARTITION BY toYYYYMM(tran_date)
ORDER BY (acc, timestamp1);

--查看分区情况

select * from system.parts where table='brch_qry_dtl';
