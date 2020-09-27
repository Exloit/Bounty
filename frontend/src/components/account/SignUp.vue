<template>
    
</template>
<script>
import RegisterAccount from './register/RegisterAccount'
import FillingInfo from './register/FillingInfo'

export default {
    components: {
        RegisterAccount,
        FillingInfo
    },
    data() {
        return {
            active: 0,
            showStepRegister: true,
            showStepBind: false,
            showStepFill: false
        }
    },
    computed: {
        showStep1 (){
            return this.showStepRegister
        },
        showStep2 (){
            return this.showStepBind
        },
        showStep3 (){
            return this.showStepFill
        }
    },
    methods: {
        next() {
            this.skip()
        },
        skip () {
            if (this.active === 0){
                this.active++
                this.showStepRegister = false
                this.showStepBind = true
                return
            }
            if (this.active++ === 1){
                this.showStepBind = false
                this.showStepFill = true
                return
            }
            this.active = 0
            this.showStepRegister = true
            this.showStepFill = false
        }
    }
}
</script>

<style scoped>
.signup-section{
    width: 666px;
    margin: 100px auto;
    background-color: #fff;
    padding: 10px 20px;
}
</style>