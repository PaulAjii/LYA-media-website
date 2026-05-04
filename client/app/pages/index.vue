<script setup lang="ts">
const { setDark } = useHeaderTheme()

onMounted(() => setDark())
</script>

<template>
    <main class="min-h-screen text-gray-700 pb-20 pt-12 bg-gray-50">
        <HeroComponent />
        <section class="w-full grid grid-cols-1 gap-12 mt-8 py-10 px-8 md:grid-cols-4 md:gap-2">
            <p class="w-full px-6 font-oswald font-black text-4xl md:pl-14 md:pr-6">
                <span class="text-blue-700">LOYOLA</span>
                <br />YOUTH ASSEMBLY
            </p>
            <div class="px-6 font-montserrat flex gap-5 flex-col md:col-span-3 md:flex-row">
                <p class="text-justify flex-1">
                    Loyola Youth Assembly, <b>LYA</b> as is usually called, is the Youth Arm of Christ Apostolic Church,
                    Oke
                    Alafia, Loyola District Headquarters.
                </p>
                <div class="flex-1">
                    <p>Our services schedules are as follows:</p>
                    <p class="pl-4 font-black">Every first, second, and fourth Sundays - 09:00 AM</p>
                    <p class="pl-4 font-black">Every third Sundays - 08:00 AM</p>
                    <p class="pl-4 font-black">Midweek Services every third fridays - 05:00 AM</p>
                    <p class="my-4">
                        LYA is located at Christ Apostolic Church, Oke Alafia, opp Nigerian Immigration
                        Services, Old Ife
                        Road, Ibadan
                    </p>
                    <p>Join us in church and experience Christ Jesus.</p>
                </div>
            </div>
        </section>
    </main>
</template>
