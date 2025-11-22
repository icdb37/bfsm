export interface ListSelecter {
  getTitle : () => string,
  getNote : () => string
};


interface SelectGetter {
  getTitle : () => string,
  getNote : () => string
};

class SelectItem {
  data : SelectGetter;
  checked : boolean;

  constructor(data : SelectGetter) {
    this.data = data;
  }

  getTitle() : string {
    return this.data.getTitle();
  }
  getNote() : string {
    return this.data.getNote();
  }
};

export class SelectList {
  items : SelectItem[];

  reset(datas : SelectGetter[], choose:number[]) {
    this.items = [];
    for (let i = 0; i < datas.length; i++) {
      this.items.push(new SelectItem(datas[i]));
    }
    for (let i = 0; i < choose.length; i++){
      this.items[choose[i]].checked = true;
    }
  }
};


class Goods implements ListSelecter {
  name : string;
  spec : string;
  size : string;

  constructor(name : string, spec : string, size : string) {
    this.name = name;
    this.spec = spec;
    this.size = size;
  }

  getTitle() : string {
    return this.name;
  }
  getNote() : string {
    return this.spec + "-" + this.size;
  }
}

export const examples : ListSelecter[] = [
  new Goods("a1", "a2", "a3"),
  new Goods("b1", "b2", "a3"),
  new Goods("c1", "b2", "a3"),
  new Goods("d1", "b2", "a3"),
  new Goods("e1", "b2", "a3"),
  new Goods("f1", "b2", "a3"),
  new Goods("g1", "b2", "a3"),
  new Goods("h1", "b2", "a3"),
  new Goods("i1", "b2", "a3"),
  new Goods("j1", "b2", "a3"),
]