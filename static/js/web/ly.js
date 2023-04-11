class Home{
    inits(){
        this.acrid()
    }
    // 热门推荐随机颜色
    acrid(){
        const colors = ["#fe2d46", "#ff7f7f", "#faa90e"]
        let acids = $('.acrid').find(".acrid-item-sort")
        for (let i = 0; i < 5; i++) {
            if (i < 3) {
                $(acids[i]).css({"backgroundColor": colors[i]})
            }
        }
    }
}

let home = new Home
home.inits()