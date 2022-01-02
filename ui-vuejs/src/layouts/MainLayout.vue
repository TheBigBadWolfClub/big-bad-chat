<template>
  <q-layout view="hHh Lpr lFr">

    <q-header elevated>
      <q-toolbar>
        <!--        <q-btn flat  @click="toggleLeftDrawer" round dense icon="menu"  aria-label="Menu"/>-->

        <q-toolbar-title>
          GoChatter
        </q-toolbar-title>

        <Avatar/>

        <q-separator dark inset style="margin: 10px" vertical/>
        <ConnButton/>
        <q-separator dark inset style="margin: 10px" vertical/>
        <q-btn aria-label="Menu" dense flat icon="menu" @click="toggleRightDrawer"/>
      </q-toolbar>
    </q-header>

    <q-footer>
      <q-toolbar class="text-grey bg-dark">
        <q-toolbar-title>Footer</q-toolbar-title>
      </q-toolbar>
    </q-footer>

    <q-drawer
      v-model="leftDrawerOpen"
      :breakpoint="700"
    >
      <ChatRooms/>
    </q-drawer>

    <q-drawer
      v-model="rightDrawerOpen"
      bordered
      class="bg-grey-1"
      show-if-above
      side="right"
    >
      <q-scroll-area class="fit">
        <q-list>
          <ConnString class="config-step"/>
          <ConfigForm class="config-step"/>
          <UserAndRoomForm class="config-step"/>
        </q-list>
      </q-scroll-area>
    </q-drawer>

    <q-page-container>
      <router-view/>
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">

import {defineComponent, ref} from 'vue'
import ConfigForm from 'components/forms/ConfigForm.vue';
import ConnButton from 'components/websocket/ConnButton.vue';
import ConnString from 'components/websocket/ConnString.vue';
import Avatar from 'components/toobar/Avatar.vue';
import UserAndRoomForm from 'components/forms/UserAndRoomForm.vue';
import ChatRooms from 'components/ChatList/ChatRooms.vue';

export default defineComponent({
  name: 'MainLayout',

  components: {
    ChatRooms,
    UserAndRoomForm,
    Avatar,
    ConnString,
    ConnButton,
    ConfigForm
  },

  setup() {
    const leftDrawerOpen = ref(true)
    const rightDrawerOpen = ref(false)

    return {
      leftDrawerOpen,
      rightDrawerOpen,
      toggleRightDrawer() {
        rightDrawerOpen.value = !rightDrawerOpen.value
      }
    }
  }
})
</script>

<style lang="sass">
.config-step
  margin-bottom: 10px

</style>
