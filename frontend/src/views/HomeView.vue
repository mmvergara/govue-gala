<script setup lang="ts">
import PostItem from "@/components/PostItem.vue";
import { ref } from "vue";
import { AxiosGet } from "@/components/AxiosInstance";
import type { PostType } from "@/types/post-types";

const posts = ref<PostType[] | null>(null);

const getPosts = async () => {
  const { data } = await AxiosGet<{ posts: PostType[] }>("/post");
  posts.value = data.posts;
};
getPosts();
</script>

<template>
  <main>
    <h1 class="text-center py-8 text-[50px]">Welcome to Govue Gala</h1>
    <section class="flex justify-center items-center flex-col gap-4">
      <PostItem v-for="post in posts" :post="post" :key="post.created_at" />
    </section>
  </main>
</template>
