This is my first dev log for the project for GoShip hackclub. The project name is Forget Me Not. This is my first time making a dev log, so forgive me for my mistakes. I'm making this dev log for the first time.

So, let me basically explain what my project is. Basically, it's a bot with predefined responses that is based in Go. It is based on Ebbinghaus' forgetting curve, which is a curve in psychology that determines how the memory is saved, recalled, reinforced, etc. This is my very basic attempt to recreate how the memory is saved in a human brain with a touch of magic of Go. 


The bot that I named is Momo. In this project, you have to teach Momo some facts or some pieces of information that you want her to remember. As you teach her the information, her brain level grows, which in turn increases her level. Your main goal is just to increase her level; there is no predefined goal, it's just a fun activity. Her memory decays over time exponentially, so you have to recall the facts every few hours or few days. As more you recall the facts, the stability of the facts, i.e., active recalling, increases, therefore the facts remain over a period of time.

You can understand it with an example. Let's say in the first iteration, she forgets a fact in two hours. You take a quiz and take a review. In the second iteration, her time to forget will double, i.e., she will forget in four hours, and so on and so forth. So basically, you have to teach her facts to increase her memory stability. This is done to show the psychological principle of active recall. By the way, the two hours and the four hours is the example. In the actual programme, the time is dynamic and depends on how much time has gone since Momo learnt the fact previously. So the time here is just shown that it will double, but it's not; it is dynamic depending upon the time you are reteaching Momo the fact.

 The most difficult part of the project that I felt was persistence and the quiz stuff. In quiz stuff, you need to compare the facts and you need to design a logic in which Momo forgets a word or a sentence and then you have to type the correct word or sentence so that Momo recalls it again and her stability of the memory increases. I did many iterations of the quiz model and still it is not that good. It might have some bugs here and there and there might be a better recommendation for the quiz model. If you have a better idea for the quiz model or anything else in the project, feel free to give me a DM on Slack, create a PR, or create an issue. I want this project to grow in sense.


By the way, as I said before, this is my first dev log and I have never written dev log. There might be some mistakes here and there because I have used Google voice typing and it's pretty bad. If you have any recommendations or a way to write dev logs, you can surely suggest me as well.



I possibly might not have included the important information that is included in a dev log. You can also tell me about the information that I should improve or what should I improve in my dev logs or in my readme files, etc.


Thanks a lot. It will be your friend, Nulldonut.


-NullDonut
