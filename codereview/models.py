from django.db import models
from django.core.exceptions import ValidationError
from django.contrib.auth import get_user_model

User = get_user_model()


class Profile(models.Model):
    user = models.ForeignKey(
        User,
        on_delete=models.CASCADE
    )
    rating = models.IntegerField(
        default=0,
        null=True,
        blank=True
    )
    pref_language = models.CharField(
        max_length=255,
        default='code'
    )

    def __str__(self):
        return self.user.username


class Codebase(models.Model):
    title = models.CharField(max_length=255)
    author = models.ForeignKey(
        User,
        on_delete=models.CASCADE,
        null=True,
        blank=True
    )
    language = models.CharField(
        max_length=255,
        default='code'
    )
    viewed_count = models.IntegerField(
        default=0,
        null=True,
        blank=True
    )

    def __str__(self):
        return self.title


class Code(models.Model):
    codebase = models.ForeignKey(
        Codebase,
        on_delete=models.CASCADE
    )
    code = models.TextField()

    def __str__(self):
        return self.code


class Comment(models.Model):
    author = models.ForeignKey(
        User,
        on_delete=models.CASCADE,
        null=True,
        blank=True
    )
    codebase = models.ForeignKey(
        Codebase,
        on_delete=models.CASCADE
    )
    review = models.TextField()

    def __str__(self):
        return self.review


class Like(models.Model):
    user = models.ForeignKey(
        User,
        on_delete=models.CASCADE
    )
    codebase = models.ForeignKey(
        Codebase,
        on_delete=models.CASCADE,
        null=True,
        blank=True
    )
    comment = models.ForeignKey(
        Comment,
        on_delete=models.CASCADE,
        null=True,
        blank=True
    )

    def __str__(self):
        like_object = 'code' if self.codebase else 'comment'
        return f'{self.user.username} liked {like_object}'

    def save(self, *args, **kwargs):
        if not self.codebase and not self.comment:
            raise ValidationError('There should be codebase OR comment like')
        else:
            super(Like, self).save(*args, **kwargs)
