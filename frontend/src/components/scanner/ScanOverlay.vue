<script setup lang="ts">
</script>

<template>

  <div class="scan-box">

    <div class="scan-line"></div>

  </div>

</template>

<style scoped>

  .scan-box{
    position:absolute;
    top:30%;
    left:5%;
    width:90%;
    height:150px;
    border:2px solid red;
    border-radius:5px;
    pointer-events:none;
    overflow:hidden;
  }

  /* moving laser line */

  .scan-line{
    position:absolute;
    left:0;
    width:100%;
    height:2px;
    background:red;

    animation: scan 4s linear infinite;
  }

  @keyframes scan{

    0%{
      top:0;
    }

    50%{
      top:147px;
    }

    100%{
      top:0;
    }

  }

</style>