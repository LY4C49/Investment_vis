<template>
    <div>
      <header>
        <h1>Data Visualization Practice "Investment"</h1>
      </header>

      <section class="container">

        <section class="itemLeft">
          <ItemPage >
            <ItemOne/>
          </ItemPage>
           <ItemPage >
            <ItemTwo/>
          </ItemPage>
        </section>


        <section class="itemCenter">
          <ItemCenter>
            <ItemFive/>
          </ItemCenter>
        </section>


        <section class="itemRight">
          <ItemPage >
            <ItemThree/>
          </ItemPage>
           <ItemPage >
            <ItemFour/>
          </ItemPage>
        </section>
      </section>
    </div>
  </template>
  
  <script>
  import ItemPage from "@/components/itemPage.vue";
  
  import ItemOne from "@/components/itemOne.vue"
  import ItemTwo from "@/components/itemTwo.vue"
  import ItemThree from "@/components/itemThree.vue"
  import ItemFour from "@/components/itemFour.vue"
  import ItemCenter from "@/components/itemCenter.vue"
  import ItemFive from "@/components/itemFive.vue"

  
  
  import {inject} from "vue"
  export default {
    components: {
      ItemPage, ItemCenter,ItemOne,ItemTwo,ItemThree,ItemFour, ItemFive
    },
  
    setup(){
      let $echarts=inject("echarts")
      let $http=inject("axios")
      console.log($echarts)
      console.log($http)
    }
  };
  </script>
  
  <style lang="less">
  header {
    height: 1rem;
    width: 100%;
    //background-color: rgba(0, 0, 255, 0.1);
    h1 {
      font-size: 0.375rem;
      color: #fff;
      text-align: center;
      line-height: 1rem;
    }
  }
  

  .container {

    min-width: 1200px;
    max-width: 2048px;
    min-height: 1200px;
    max-height: 1980px;
    margin: 0 auto;
    padding: 0.125rem 0.125rem 0;
    // background-color: gray (DEVELOP ONLY);
    display: flex;

    .itemLeft{
      flex: 7
    }
    .itemRight {
      flex: 5;
    }
    .itemCenter {
      flex: 5;
      height: 10.5rem;
      //border: 1px solid blue;
      padding:  0.125rem;
    }
  }
  </style>