<template>
    <view class="form-container">
        <uni-forms :modelValue="batch">
            <uni-forms-item label="采购名称" name="name">
                <uni-easyinput v-model="batch.name" placeholder="请输入名称" />
            </uni-forms-item>
            <uni-forms-item label="采购描述" name="desc">
                <uni-easyinput v-model="batch.desc" placeholder="请输入描述" />
            </uni-forms-item>
            <uni-forms-item label="企业商品" name="companies">
                <view class="append-item-container">
                    <button size="mini" type="primary" @click="onAddCompany">+</button>
                </view>
                <uni-list>
                    <view class="append-item-container" v-for="(item,index) in batch.companies">
                        <uni-list-item :key="index" :title="item.company.company_name" :note="''+item.amount_total" />
                        <button class="append-item-chg" type="success" size="mini"
                            @click="onChgCompany(index)">编辑</button>
                        <button class="append-item-del" type="warn" size="mini" @click="onDelCompany(index)">删除</button>
                    </view>
                </uni-list>
            </uni-forms-item>
            <uni-forms-item label="额外费用" name="extras">
                <view class="append-item-container">
                    <button size="mini" type="primary" @click="onAddExpense">+</button>
                </view>
                <uni-list>
                    <view class="append-item-container" v-for="(item,index) in batch.extras">
                        <uni-list-item :key="index" :title="item.name" :note="''+item.amount" />
                        <button class="append-item-chg" type="success" size="mini"
                            @click="onChgExpense(index)">编辑</button>
                        <button class="append-item-del" type="warn" size="mini" @click="onDelExpense(index)">删除</button>
                    </view>
                </uni-list>
            </uni-forms-item>
        </uni-forms>

        <view class="actions">
            <button type="primary" @click="submit">保存</button>
            <button type="default" @click="cancel">取消</button>
        </view>
        <view>
            <uni-popup ref="popupTip" type="message">
                <uni-popup-message :type="tip.type" :message="tip.message" :duration="2000"></uni-popup-message>
            </uni-popup>
        </view>

        <!-- 商品列表 -->
        <view>
            <uni-popup ref="popupCompnay" type="dialog">
                <uni-popup-dialog title="采购企业信息" type="info" :before-close="true" @confirm="onSubmitCompany"
                    @close="onCancelCompany">
                    <company-commodity v-model:batch="batch" v-model:company="company" v-model:mode="mode" />
                </uni-popup-dialog>
            </uni-popup>
        </view>
        <!-- 费用列表 -->
        <view>
            <uni-popup ref="popupExpense" type="dialog">
                <uni-popup-dialog model="input" type="info" :before-close="true" cancelText="取消" confirmText="确认"
                    title="额外费用" @confirm="onAddExpenseConfirm" @close="onAddExpenseCancel">
                    <purchase-expense-extra v-model:batch="batch" v-model:expense="expense" />
                </uni-popup-dialog>
            </uni-popup>
        </view>
    </view>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import { BaseURL, Commodity } from '@/xapi/xapi';
    import { onLoad } from '@dcloudio/uni-app';
    import { PurchaseBatch, PurchaseCompany, PurchaseExpense } from '@/xapi/model/purchase';

    const batch = ref<PurchaseBatch>(new PurchaseBatch());
    const company = ref<PurchaseCompany>(new PurchaseCompany());
    const mode = ref(modeCreate);
    const modeCreate = 'create';
    const modeUpdate = 'update';
    const expense = ref<PurchaseExpense>(new PurchaseExpense());
    let expenseIndex = -1;
    let purchaseURL = `${BaseURL}/api/v1/purchase`; // 页面加装时初始化
    let companyURL = `${BaseURL}/api/v1/company`; // 页面加装时初始化
    let batchID = '';

    const popupTip = ref(null);
    const tip = ref({
        type: 'success',
        message: '成功'
    })
    const popupCompnay = ref(null);
    const popupExpense = ref(null);

    function onAddCompany() {
        mode.value = modeCreate;
        popupCompnay.value.open();
    }
    function onSubmitCompany() {
        let i = 0
        for (; i < batch.value.companies.length; i++) {
            if (batch.value.companies[i].company.company_id == company.value.company.company_id) {
                break;
            }
        }
        if (i < batch.value.companies.length) {
            batch.value.companies[i] = { ...company.value };
        } else {
            batch.value.companies.push({ ...company.value });
        }
        console.log("onSubmitCommpany", JSON.stringify(company.value));
        company.value.reset();
        popupCompnay.value.close();
    }
    function onCancelCompany() {
        company.value.reset();
        popupCompnay.value.close();
    }
    function onChgCompany(pos : any) {
        company.value = batch.value.companies[pos];
        console.log(JSON.stringify(company.value))
        console.log(JSON.stringify(batch.value))
        popupCompnay.value.open();
    }
    function onDelCompany(pos : any) {
        batch.value.companies.splice(pos, 1);
    }

    function onAddExpense() {
        expenseIndex = -1;
        expense.value = new PurchaseExpense();
        popupExpense.value.open();
    }
    function onDelExpense(pos : any) {
        batch.value.extras.splice(pos, 1)
    }

    function onChgExpense(pos : any) {
        expenseIndex = pos;
        expense.value = batch.value.extras[pos];
        popupExpense.value.open();
    }

    function onAddExpenseConfirm() {
        if (expenseIndex == -1) {
            batch.value.extras.push({ ...expense.value });
        }
        popupExpense.value.close();
    }

    function onAddExpenseCancel() {
        popupExpense.value.close();
    }

    function submit() {
        uni.request({
            url: purchaseURL + "/batch/" + batchID,
            method: 'PUT',
            data: batch.value,
            success: (res) => {
                console.log("success", res)
                if (res.statusCode == 200) {
                    tip.value.type = 'success'
                    tip.value.message = '成功'
                    setTimeout(
                        () => {
                            uni.navigateBack();
                            uni.$emit('purchaseBatchRefresh')
                        },
                        500);
                } else {
                    tip.value.type = 'error'
                    tip.value.message = res.data.message
                }
                popupTip.value.open();
            }
        })
    }
    function getDetail(id : string) {
        uni.request({
            url: purchaseURL + "/batch/" + id,
            method: 'GET',
            success: (res) => {
                batch.value = { ...res.data };
            }
        })
    }

    function cancel() {
        uni.navigateBack();
    }
    onLoad((options) => {
        purchaseURL = `${BaseURL}/api/v1/purchase`;
        companyURL = `${BaseURL}/api/v1/company`;
        batchID = "4525b2ac-5b04-4b1c-9ab9-a9a091d7f1b6";
        if (options.id) {
            batchID = options.id;
        }
        getDetail(batchID);
    })
</script>

<style>
    .form-container {
        padding: 16rpx;
        width: 600rpx;
    }

    .append-item-container {
        display: flex;
        align-items: center;
    }

    .append-item-chg {
        margin-right: 0px;
    }

    .append-item-del {
        margin-right: 0px;
        margin-left: 0px;
    }

    .actions {
        display: flex;
        justify-content: flex-end;
        gap: 12rpx;
        margin-top: 16rpx;
    }

    .company-commodity-container {
        /* 限制高度 */
    }

    .company-commodity-container-item {
        display: flex;
        align-items: center;
    }

    .company-commodity-container-count {
        width: 10%;
    }
</style>