<script setup lang="ts">
let err = "";
const { data, pending, error } = await useFetch<{ posts: any }>(
  `http://localhost:8080/post`,
  {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    cache: "no-cache",
  }
);
if (pending.value) {
  err = "Loading...";
} else if (error.value) {
  console.log(error.value);
} else {
  err = "";
}
</script>

<template>
  <main>
    <p v-show="err">{{ err }}</p>
    <PostItem v-for="post in data?.posts" :post="post" :key="post.created_at" />
  </main>
</template>
