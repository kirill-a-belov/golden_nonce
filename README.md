Bitcoin mining research: The best Golden Nonce finding strategy
---------------------------------------------------------------
**What is Bitcoin Nonce?**

The essence of the Bitcoin block mining is the finding of such hash sum of the new block, which will be less than global 
target number (complexity of mining). Since we can't change the contents of the prepared block (actually - we can, but 
in this case our mining will be very long time) - we have to add special number to block's data. 
Exactly this number (Nonce) defines which hash sum we will get. The Nonce who lead us to getting proper hash getting 
is Golden Nonce. I was interested if there is a difference between the way you pick another nonce for hash calculating. 
I selected two different strategy for Nonce picking and try to find what will be better: go over from zero till the 
maximum in a row or randomly picking numbers in this interval. What strategy will allow you to find your 
"Golden Nonce" first?

**Simulation**

My simulation has written in Go (Golang). In the frame of it i took some nonce maximum sizes and two numbers (1 and 100) 
of iterations for finding target Nonce for each size. For each Nonce size i also generated new pseudorandom numbers 
sequence (i used "map" data structure for access time reducing).

>First of all - would like to say, this result is only theoretical because of each miner find Nonce for his own block 
>with different set of transactions and data. But lets image we live in ideal world and all miners started to mine at 
>the same time on a same hardware and software and using equal blocks, the difference only in Nonce finding strategy!


**Graphs**

Here is a graphs which shows average percentage of ratio of attempts to the maximum Nonce size.
In other words it show how many attempts computer needed to find target Nonce in a range from zero to maximum Nonce 
size for each strategy. Lower value is better.

![\[graph\]](https://github.com/kirill-a-belov/golden_nonce/1itr.jpeg)
One iteration.

![\[graph\]](https://github.com/kirill-a-belov/golden_nonce/100itr.jpeg)
Hundred iterations.

As we can see for 1 iteration GoOver strategy is obviously better than Picking. But for 100 iterations - they show 
practically identical results. And Bitcoin nonce is 32-bit (4-byte) field. It means that it can has maximum size 
is 4.2 Billion. And you certanly wil have to do more than one hundred attempts before your can find Golden Nonce.

**Conclusions:**

**1**. On a big numbers with big iteration count - it doesn't matter what Nonce finding strategy you'll pick.

**2**. This results can be explain by normal distribution ([Normal distribution](https://en.wikipedia.org/wiki/Normal_distribution)) 
which allow us to say that on big numbers, success probabilities of our strategies will be approximately equal.

**3**. If you really want to find strategy with different results you should use different types or random number 
generators (non-compliant to normal distribution) , based for example on today's horoscope. But probability to mine 
block before others will be like probability of winning in lottery. Statistic - a harsh thing!