<template>
  <main>
    <p v-show="!!error">{{ error }}</p>
    <PostItem v-for="post in data?.posts" :post="post" :key="post.created_at" />
  </main>
</template>

<script setup lang="ts">
import PostItem from "@/components/post-item.vue";
import type { PostType } from "@/types/post-types";
import { getRequest } from "~/lib/fetching";
const { data, error } = await getRequest<{ posts: PostType[] }>("/post");
</script>
