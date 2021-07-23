<template>
  <div class="main">
    <img :src="`${$root.API_URL}/file?path=${files[fileId]?.path}`" alt="" draggable="false" />
    <div @click="prev()" class="clickable prev">
      <img src="../asset/left.svg" alt="" draggable="false" />
    </div>
    <div @click="next()" class="clickable next">
      <img src="../asset/right.svg" alt="" draggable="false" />
    </div>
    <div class="total">{{ fileId + 1 }} / {{ files.length }}</div>
    <div class="name">{{ files[fileId]?.path.split('/').pop() }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: {},
  async mounted() {
    const info = await RestApi.file.getInfo();

    this.files = info.files;
    this.fileId = 0;
    this.file = info.files[0].path;
    if (this.file) {
      this.fileId = this.files.findIndex((x: any) => x.path === this.file);
    }
  },
  methods: {
    next() {
      this.fileId += 1;
      if (this.fileId > this.files.length - 1) {
        this.fileId = this.files.length - 1;
      }
    },
    prev() {
      this.fileId -= 1;
      if (this.fileId <= 0) {
        this.fileId = 0;
      }
    },
  },
  data: () => {
    return {
      files: [],
      file: '',
      fileId: 0,
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  display: flex;
  height: 100%;
  user-select: none;

  img {
    width: 100%;
    // height: 100%;
    display: flex;
    margin: auto;
  }

  .prev {
    position: fixed;
    left: 10px;
    top: 50%;
  }

  .next {
    position: fixed;
    right: 10px;
    top: 50%;
  }

  .next,
  .prev {
    background: rgba(0, 0, 0, 0.45);
    padding: 7px 10px;
    border-radius: 4px;
  }

  .total {
    background: rgba(0, 0, 0, 0.45);
    position: fixed;
    left: 50%;
    bottom: 10px;
    padding: 5px 10px;
    border-radius: 4px;
    font-size: 14px;
    color: #fefefe;
    transform: translate(-50%, -50%);
  }

  .name {
    background: rgba(0, 0, 0, 0.45);
    position: fixed;
    top: 30px;
    color: #fefefe;
    padding: 5px 10px;
    border-radius: 4px;
    font-size: 14px;
    left: 50%;
    transform: translate(-50%, -50%);
    white-space: nowrap;
  }
}
</style>
