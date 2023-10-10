<template>
  <div>{{ $route.params.postid }}</div>
</template>

<script setup lang="ts">
// get postid param
import { useRoute } from "vue-router";
const route = useRoute();
console.log(route.params.postid);
</script>
