export class RefCompany {
    company_id : string;
    company_name : string;
    
    constructor(){
        this.company_id = "";
        this.company_name = "";
    }
}

// PurchaseGoods 采购商品
export class PurchaseGoods {
    id : string;      // 商品ID
    hash : string;     // 哈希
    name : string;     // 名称
    desc : string;     // 描述
    sepc : string;     // 规格
    size : string;     // 尺寸
    validity : number; // 有效期
    price : number;    // 价格
    count : number;     // 数量
    amount : number;    // 金额
    produced_at : string; // 生产日期
    expired_at : string; // 过期日期
    storeage : string; // 存储位置
    
    constructor(arg:PurchaseGoods) {
        this.id = arg.id;
        this.hash = arg.hash;
        this.name = arg.name;
        this.desc = arg.desc;
        this.sepc = arg.sepc;
        this.size = arg.size;
        this.validity = arg.validity;
        this.price = arg.price;
        this.count = arg.count;
        this.amount = arg.amount;
        this.produced_at = arg.produced_at;
        this.expired_at = arg.expired_at;
        this.storeage = arg.storeage;
        if (this.count == undefined || this.count == null) {
            this.count = 0;
        }
        if (this.price == undefined || this.price == null){
            this.price = 0;
        }
        if (this.amount == undefined || this.amount == null){
            this.amount = 0;
        }
    }
}

// PurchaseCompany 采购企业
export class PurchaseCompany {
    company : RefCompany; // 企业基本信息
    desc : string;        // 企业采购描述
    amount_extra : number; // 额外金额
    amount_goods : number; // 商品金额
    amount_total : number; // 总金额
    goods : PurchaseGoods[];  //企业商品列表
    
    constructor(){
        this.company = new RefCompany();
        this.goods = [];
        if (this.amount_extra == undefined || this.amount_extra == null) {
            this.amount_extra = 0;
        }
        if (this.amount_goods == undefined || this.amount_goods == null) {
            this.amount_goods = 0;
        }
        if (this.amount_total == undefined || this.amount_total == null) {
            this.amount_total = 0;
        }
    }
    reset(){
        this.amount_extra = 0;
        this.amount_goods = 0;
        this.amount_total = 0;
        this.goods = [];
        this.company = new RefCompany();
        this.desc = "";
    }
}

// PurchaseExpense 额外费用
export class PurchaseExpense{
    name: string;
    desc: string;
    amount: number;
    
    constructor() {
        this.name = '';
        this.desc = '';
        this.amount = 0;
    }
    reset(){
        this.name = '';
        this.desc = '';
        this.amount = 0;
    }
}

// PurchaseBatch 采购批次
export class PurchaseBatch {
    id : string;   // 批次ID
    name : string; // 批次名称
    desc : string; // 批次描述
    status: number; // 批次状态 
    created_at : string; // 创建时间
    amount_goods: number; // 商品金额
    amount_extra: number; // 额外金额
    amount_total: number; // 总金额
    extras: PurchaseExpense[]; // 额外费用
    companies : PurchaseCompany[]; // 采购企业列表
    
    constructor(){
        this.companies = [];
        this.extras = [];
        if (this.amount_extra == undefined || this.amount_extra == null){
            this.amount_extra = 0;
        }
        if (this.amount_goods == undefined || this.amount_goods == null){
            this.amount_goods = 0;
        }
        if (this.amount_total == undefined || this.amount_total == null){
            this.amount_total = 0;
        }
    }
}


// PurchaseGoods 采购商品
export class PurchaseGoodsX {
    purchase_id: string;   // 采购批次标识
    purchase_name: string; // 采购批次名称
    company_id : string;   // 企业标识
    company_name : string; // 企业名称
    id : string;      // 商品ID
    hash : string;     // 哈希
    name : string;     // 名称
    desc : string;     // 描述
    sepc : string;     // 规格
    size : string;     // 尺寸
    validity : number; // 有效期
    price : number;    // 价格
    count : number;     // 数量
    amount : number;    // 金额
    created_at: string; // 创建时间
    produced_at : string; // 生产日期
    expired_at : string; // 过期日期
    storeage : string; // 存储位置
    
    constructor(v:any) {
        if(v != undefined || v != null) {
            this.set(v);
            return;
        }
        this.id = "";
        this.hash = "";
        this.name = "";
        this.desc = "";
        this.sepc = "";
        this.size = "";
        this.validity = 0;
        this.price = 0;
        this.count = 0;
        this.amount = 0;
        this.created_at = "";
        this.produced_at = "";
        this.expired_at = "";
        this.storeage = "";
        this.company_id = "";
        this.company_name = "";
        this.purchase_id = "";
        this.purchase_name = "";
    }
    set(arg:PurchaseGoodsX) {
        this.id = arg.id;
        this.hash = arg.hash;
        this.name = arg.name;
        this.desc = arg.desc;
        this.sepc = arg.sepc;
        this.size = arg.size;
        this.validity = arg.validity;
        this.price = arg.price;
        this.count = arg.count;
        this.amount = arg.amount;
        this.created_at = arg.created_at;
        this.produced_at = arg.produced_at;
        this.expired_at = arg.expired_at;
        this.storeage = arg.storeage;
        this.company_id = arg.company_id;
        this.company_name = arg.company_name;
        this.purchase_id = arg.purchase_id;
        this.purchase_name = arg.purchase_name;
        if (this.count == undefined || this.count == null) {
            this.count = 0;
        }
        if (this.price == undefined || this.price == null){
            this.price = 0;
        }
        if (this.amount == undefined || this.amount == null){
            this.amount = 0;
        }
        if (this.id == undefined || this.id == null) {
            this.id = "";
        }
    }
}