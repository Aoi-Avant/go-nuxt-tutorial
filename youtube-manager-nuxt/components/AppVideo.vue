<template>
  <nuxt-link :to="'/video/${videoId}'">
    <div class="box video-box">
      <article class="media">
        <div class="media-left">
          <figure class="image is-220x120">
            <img :src="item.snippet.thumbnails.medium.url" alt="Image" />
          </figure>
        </div>
        <div class="media-content">
          <div class="content">
            <p>
              <strong>
                {{ item.snippet.title }}
              </strong>
              <br />
              <small>
                {{ item.snippet.publishedAt }}
              </small>
              <small>
                {{ item.snippet.channelTitle }}
              </small>
              <br />
              {{ item.snippet.description | omit }}
            </p>
          </div>
        </div>
      </article>
    </div>
  </nuxt-link>
</template>

<script>
export default {
  filters: {
    omit: value => {
      if (!value) {
        return ''
      }

      if (value.length > 120) {
        return value.substr(0, 120) + '...'
      }

      return value
    },
  },
  props: {
    item: {
      type: Object,
      default: null,
    },
    videoId: {
      type: String,
      default: '',
    },
  },
}
</script>
