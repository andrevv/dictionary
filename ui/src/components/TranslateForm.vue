<template>
	<div class="mb-3">
		<input type="text" class="form-control" id="text" placeholder="Enter a word to translate..." v-model="text"/>
	</div>
	<div class="mb-3">
		<button type="button" class="btn btn-primary" @click.prevent="translate">Translate</button>
	</div>
	<div class="mb-3">
		<ul v-for="translation in translations" :key="translation">
			<li>{{translation}}</li>
		</ul>
	</div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
export default defineComponent({
	name: 'TranslateForm',
	setup() {
		const text = ref('time')
		const translations = ref()

		const translate = async () => {
			const resp = await fetch(`http://localhost:8080/api/translate?lang=en-ru&text=${text.value}`)
			translations.value = await resp.json()
		}

		return {
			text,
			translations,
			translate			
		}
	}
})
</script>